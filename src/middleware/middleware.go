package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	v "centralize-authentication-go/src/version"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func ResponseBuilder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a new custom response writer
		w := &ResponseWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		// Set the custom writer to the context
		ctx.Writer = w

		// Process the request
		ctx.Next()

		// If the response status is OK, modify the body
		if ctx.Writer.Status() == http.StatusOK {
			// Try to parse the captured response body as a JSON object
			responseBody := make(map[string]interface{})
			if err := json.Unmarshal([]byte(w.body.String()), &responseBody); err == nil {
				// Move the original data into the "data" field and add "meta"
				newResponse := make(map[string]interface{})
				newResponse["meta"] = map[string]interface{}{"version": v.AppVersion}
				newResponse["data"] = responseBody

				responseFinal, _ := json.Marshal(newResponse)

				// Write the new modified response with "meta" and "data"
				w.ResponseWriter.WriteString(string(responseFinal))
				w.body.Reset()
			} else {
				// In case unmarshalling fails, fallback to raw content with "meta"
				ctx.JSON(http.StatusOK, gin.H{
					"meta": gin.H{"version": "v1.0.0"},
					"data": w.body.String(), // Include original raw response body
				})
			}
		}
	}
}

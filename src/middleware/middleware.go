package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"centralize-authentication-go/src/middleware/model"
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

		if ctx.Writer.Status() == http.StatusOK {
			// Try to parse the captured response body as a JSON object
			responseSuccess := model.ResponseSuccess{
				Meta: model.MetaVersion{
					Version: v.AppVersion,
				},
			}
			responseBody := make(map[string]interface{})
			if err := json.Unmarshal([]byte(w.body.String()), &responseBody); err == nil {
				responseSuccess.Data = responseBody
				responseFinal, _ := json.Marshal(responseSuccess)
				w.ResponseWriter.WriteString(string(responseFinal))
				w.body.Reset()
			}
		} else {
			responseError := model.ResponseError{
				Meta: model.MetaVersion{
					Version: v.AppVersion,
				},
			}
			responseBody := make(map[string]interface{})
			if err := json.Unmarshal([]byte(w.body.String()), &responseBody); err == nil {
				responseError.Reason = responseBody["reason"].(string)
				responseFinal, _ := json.Marshal(responseError)
				w.ResponseWriter.WriteString(string(responseFinal))
				w.body.Reset()
			}
		}
	}
}

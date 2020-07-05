package importer

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMakeOpenAPI3Importer(t *testing.T) {
	logger := logrus.New()
	basePath := ""
	importer := MakeOpenAPI3Importer(logger, basePath, "")
	t.Log(importer)
}

func TestLoadOpenAPI3(t *testing.T) {
	spec := `openapi: "3.0"
info:
  title: Simple
paths:
  /test/:
    get:
      responses:
        200:
          description: "200 OK"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/SimpleObj"
        500:
          $ref: "#/components/responses/500Response"
components:
  schemas:
    SimpleObj:
      type: object
      properties:
        name:
          type: string
    SimpleObj2:
      type: object
      properties:
        name:
          type: SimpleObj
  responses:
    500Response:
      description: Internal Server Error
      schema:
        $ref: "#/components/schemas/SimpleObj"
`
	logger := logrus.New()
	basePath := ""
	importer := MakeOpenAPI3Importer(logger, basePath, "")
	result, err := importer.Load(spec)
	assert.NoError(t, err)
	t.Log(result)
}

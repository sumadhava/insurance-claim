{
    "swagger": "2.0",
    "info": {
        "description": "Swagger Insurance Claim Application.",
        "version": "1.0.0",
        "title": "Swagger Insurance Claim Application",
        
        "contact": {
            "email": ""
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        }
    },
    "host": "localhost:3000",
    "basePath": "/tfbc",
    
    "schemes": [
        "http"
    ],
    "paths": {
        "/requestClaim": {
            "post": {
                
                "summary": "Request Claim",
                "description": "Request Claim",
                "operationId": "requestClaim",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "LC object",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LC"
                        }
                    }
                ],
                "responses": {
                    "405": {
                        "description": "Invalid input"
                    }
                }
            }
        },
        "/issueLC": {
            "post": {
                
                "summary": "Issue LC",
                "description": "Issue LC",
                "operationId": "issueLC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "LC ID object",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LCID"
                        }
                    }
                ],
                "responses": {
                    "405": {
                        "description": "Invalid input"
                    }
                }
            }
        },
        "/acceptLC": {
            "post": {
                
                "summary": "Accept LC",
                "description": "Accept LC",
                "operationId": "acceptLC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "LC ID object",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LCID"
                        }
                    }
                ],
                "responses": {
                    "405": {
                        "description": "Invalid input"
                    }
                }
            }
        },
        "/getLC": {
            "post": {
                
                "summary": "Get LC",
                "description": "Get LC",
                "operationId": "getLC",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "LC object",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LCID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "type": "object",
                            "items": {
                                "$ref": "#/definitions/LC"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid status value"
                    }
                }
            }
        },
        "/getLCHistory": {
            "post": {
                
                "summary": "Get LC History",
                "description": "Get LC History",
                "operationId": "getLCHistory",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "in": "body",
                        "name": "body",
                        "description": "LC object",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/LCID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/LC"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid status value"
                    }
                }
            }
        }
    },
    "definitions": {
        "LC": {
            "type": "object",
            "properties": {
                "ClaimID": {
                    "type": "string",
                    "description": "Claim ID"
                },
                "PolicyNumber": {
                    "type": "string",
                    "description": "Policy Number"
                },
                "EntryDate": {
                    "type": "string",
                    "description": "Entry Date"
                },
                "InsuranceCompany": {
                    "type": "string",
                    "description": "Insurance Company"
                },
                "PlaceOfService": {
                    "type": "string",
                    "description": "Place Of Service"
                },
                "ProviderName": {
                    "type": "string",
                    "description": "Provider Name"
                },
                "ClaimAmount": {
                    "type": "string",
                    "description": "Claim Amount"
                },
                "DateOfService": {
                    "type": "string",
                    "description": "Date Of Service"
                },
                "DiagnosCode": {
                    "type": "string",
                    "description": "Diagnos Code"
                },
                "Procedure Code": {
                    "type": "string",
                    "description": "Procedure Code"
                },
                "TypeOfService": {
                    "type": "string",
                    "description": "Type Of Service"
                }

            }
        },
        "LCID": {
            "type": "object",
            "properties": {
                "ClaimID": {
                    "type": "string",
                    "description": "Claim ID"
                }
            }
        }
    },
    "externalDocs": {
        "description": "Find out more about Swagger",
        "url": "http://swagger.io"
    }
}

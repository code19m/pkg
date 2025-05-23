# API Schema Conventions

This document outlines our API design principles and standardized response formats to ensure consistency across all API endpoints.

## URL Structure

- Follow REST conventions as outlined in [REST API Resource Naming Guide](https://restfulapi.net/resource-naming/)
- Avoid nesting resources in URLs when possible

## Request Conventions

- Use `snake_case` for all JSON property names
- Standard pagination parameters:
  - `limit`: Number of items to return
  - `offset`: Number of items to skip

## Response Formats

### Paginated List Response

Use this format when returning a paginated list of resources:

```json
{
  "count": 100, // Total count of available items
  "page_list": [
    // Array of items for the current page
    {
      // Individual item schema
    }
  ]
}
```

### Non-Paginated List Response

Use this format when returning a complete (non-paginated) list of resources:

```json
{
  "list": [
    // Array of all items
    {
      // Individual item schema
    }
  ]
}
```

### Error Response

All error responses should follow this standard format:

**Headers:**

- `x-request-id`: Unique identifier for the request (required for all responses)

**Body:**

```json
{
  "error": {
    "code": "string", // Error code (can be empty)
    "message": "string", // Human-readable error message for debugging
    "trace": "string", // Trace ID for debugging (hidden in production)
    "fields": {
      // Field-specific errors for form validation
      // map[string]string of field names to error messages
    },
    "details": {
      // Additional error details (hidden in production)
      // map[string]string with detailed error information
    }
  }
}
```

> **Note for Frontend Development:** Focus on the `error.code` field of the body and `x-request-id` header. Other fields are primarily for debugging purposes,
> and should not be relied upon for user-facing error messages.

## API Versioning

Include API version in the URL path:

```
/api/v1/resources
```

---

These conventions ensure consistent API design and make it easier to integrate with our services. Always follow these guidelines when designing or implementing new API endpoints.

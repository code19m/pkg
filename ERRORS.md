# Error Handling Conventions

This document outlines our error handling standards for both frontend and backend development. These conventions ensure consistent error handling across our applications.

## Frontend Error Handling Rules

1. **UI Validation Errors**: Display error messages directly on the UI without dealing with backend.

2. **Server Response Error Handling**:

   - This error handling mechanism doesn't apply to authentication errors (401) - these should be handled separately in interceptors.
   - For unknown error codes (even if status code is 4xx) or server errors (5xx): Display generic error message (e.g. Something went wrong) with **error code** and **request id** if available.
   - For client validation/not_found/conflict errors (4xx) with known error codes: Display translated meaningful error messages.
   - Consider that not all client error codes should be shown to the user. For example, if frontend code sends wrong input data to the backend due to a bug and the backend returns an error code like "USER_NOT_FOUND", then this error should be considered an unknown error code for the user and should be handled as a generic error.
   - Consider localization of error messages for both generic and known errors.

## Backend Error Handling Rules

### Error Types

1. **Input Validation Errors**: Defined at controller layer.
2. **Business Logic Validation/Conflict Errors**: Defined at use case layer.
3. **Internal Server Errors**: All other errors (by default).

### Implementation Rules

1. **Use the `errx` package** for all error handling - it provides typing, coding, and debugging features.

2. **Define error codes in domain layer**.

3. **Controller Layer**:

   - Use `playground/validator/v10` for input validation.
   - Return `errx` errors with validation type. `errx.WithType(errx.T_Validation)`
   - Use `errx.WithFields()` option to add validation messages by fields.
   - TODO: Add reusable validator package to common pkg repo.

4. **Use Case Layer**:

   - Handle business logic validation based on downstream error codes.
   - Apply appropriate error types with `errx.WrapWithTypeOnCodes()`.
   - Check for ownership violations or other business rules and return new errors with `errx.New()`.
   - Examples:
     - For "USER_NOT_FOUND" (if user_id came from user input): `errx.WrapWithTypeOnCodes(err, errx.T_NotFound, "USER_NOT_FOUND")`
     - For "USER_CONFLICT" (if create input params came from user input): `errx.WrapWithTypeOnCodes(err, errx.T_Conflict, "USER_CONFLICT")`
     - If user.organization_id != resource.organization_id: `errx.New("This organization doesn't belong to this user", errx.WithType(errx.T_Forbidden), errx.WithCode("OWNERSHIP_VIOLATION"))`
     - Consider that in these examples codes are given in hardcoded strings, but in production code they should be defined in domain layer as constants.

5. **Downstream Layer Rules (service, repository, domain)**:

   - Return **ONLY** `errx` errors with `INTERNAL` type (default).
   - Error types should only be defined at use case or controller layer.

6. **Repository Layer**:
   - Add detailed error information using `errx.WithDetails()` for debugging.
   - Never include sensitive data in error details.

## Error Response Schema

- For response schema of errors, refer to `API_SCHEMAS.md` file.

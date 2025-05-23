For response schema of errors refer to 

frontend

If error is UI validation error (validation errors that are not sent to backend), then show error message on UI.

If error is server error:
    - If error kind is client error (i.e. 400, 409, 404) and error code is well known error code (e.g. 'user_already_exists'), then show error code translation message on UI. Consider that not all client error codes are well known error codes.
    - If error code is not well known error code or error kind is server (i.e. >= 500), then show generic error message (e.g. Something went wront) with error code (if exists) and request_id (if exists) on UI.

Consider localization of generic or well known error messages.


backend
There are 3 types of errors:
    1. Input validation errors
    2. Business logic validation or conflict errors
    3. Internal Server errors

Input validation errors are defined at controller layer.
Business logic validation or conflict errors are defined at use case layer.
All other errors are considered as internal server errors.

For comprehensive error handling use errx package which provides a way to set types and codes for errors and some additional abstraction for debugging (e.g. error_trace, validation fields, error details).

Define error codes in domain layer.

In controller layer:
For input validation use playground/validator/v10 package. And return errx error with validation type and validation fields. Use errx.WithFields method to add validation fields.  TODO: add reusable validator package to common pkg repo.

In use case layer:
For define business logic validation or conflict errors base on error code of downstream layer errors. For example if from userRepo.GetById() we get "USER_NOT_FOUND" error code, and id was passed from client, then we can use errx.WrapWithTypeOnCodes(err, errx.T_Validation, "USER_NOT_FOUND") to properly handle this error.
Or from repo.Create() we get "USER_CONFLICT" error code, then we can use errx.WrapWithTypeOnCodes(err, errx.T_Conflict, "USER_CONFLICT") to properly handle this error.

All downstream errors after use case layer: service, repository, domain, etc. should return ONLY errx errors with INTERNAL TYPE (which is default type for errx errors)!!! Types are defined only on use case! Or before use case at controller layer for input validation.
In repository layer:
    Add error information to errorx details as much as possible for ease of debugging, use errx.WithDetails option. Consider not to add sensitive data to details.

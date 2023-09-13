# Validator

translation based validator.

## Requirements

### Error Response

Each validation method generate `ErrorResponse`.

**Note:** ErrorResponse implements json Marshaller and toString interface by default.

#### Usage

ErrorResponse interface contains following methods:

##### AddError

Add new error to errors list.

```go
// Signature:
AddError(field, tag, message string)

// Example:
errs.AddError("username", "pattern", "Username can contains alphnum characters only!")
```

##### HasError

Check if response has error.

```go
// Signature:
HasError() bool

// Example:
if errs.HasError() {
    // Return 422 status
}else{
    // Save record and return 200 status
}
```

##### Failed

Check if field has error.

```go
// Signature:
Failed(field string) bool

// Example:
if errs.Failed("username"){
    // Username field has some error
}
```

##### FailedOn

Check if field has special error.

```go
// Signature:
FailedOn(field, err string) bool

// Example:
if errs.FailedOn("username", "pattern"){
    // Username field has pattern error
}
```

##### Errors

Get errors list as map.

```go
// Signature:
Errors() map[string]map[string]string
```

##### String

Convert error response to string.

```go
// Signature:
String() string
```

##### Messages

Get error messages only without error key

```go
// Signature:
Messages() map[string][]string
```

##### Rules

Get error rules only without error message.

```go
// Signature:
Rules() map[string][]string
```

##### MarshalJSON

Convert error response to json.

```go
// Signature:
MarshalJSON() ([]byte, error)

// Example:
if resp, err := errs.MarshalJSON(); err == nil{
    fmt.Println(string(resp)) // { "username": { "required":"username is required", "pattern":"username pattern is wrong!" } }
}
```

### Helpers

#### Invalidate

Generate invalid state for field.

```go
// Signature:
Invalidate(field, err string) ErrorResponse
```

### Validation functions

You can access custom validations by function.

```go
import "github.com/gomig/validator"
if validator.IsIDNumber(1234) {
    // This is a valid id number
}
```

### Field Name

Validator get field friendly name from struct tag or ValidatorParam.

For struct tags, validator get field name from `field`, `json`, `form` and `xml` tag in order. if field name not specified by tag validator use Field Struct name as field name.

For ValidatorParam, validator get field name from `ValidatorParam.Name` field.

```go
type Person struct{
    Name string `validate:"required" json:"firstname"`
    Family string `validate:"required"`
}

// => error response
// {
//     "firstname":{...},
//     "Family":{...}
// }
```

### Helper Tags

Validator use three special tags for translating error message.

#### vTitle

This tag use as field title in error messages. you can combine this tag with locale key for use translation. if this parameter not passed validator use field name as title.

```go
type Person struct{
    Name string `validate:"required" vTitle_fa:"نام" vTitle:"Firstname"`
}
```

#### vParam

When use validator that contains another field (like eqcsfield), you can use this tag for parameter title like `vTitle` tag. if this parameter not passed validator use validation param itself.

```go
type Person struct{
    Password string `validate:"required" form:"pswd" vTitle_fa:"رمز"`
    PasswordRe string `validate:"eqcsfield:password" form:"pswd_re" vTitle_fa:"تایید رمز" vTitle:"Password repeat" vParam_fa:"رمز" vParam:"password"`
}
```

#### vFormat

When use this tag validator parameter formatted as number.

```go
type Transaction struct{
    Amount string `validate:"gt=1000" xml:"amount" vTitle_fa:"مبلغ" vFormat`
}

//  => Error Message: amount must be greater than 1,000
```

## Create New Validator

Validator based on `github.com/go-playground/validator/v10` and `github.com/gomig/translator` packages.

**Note:** Validator library contains translation for go builtin validator functions.

```go
// Signature:
NewValidator(t translator.Translator, locale string) Validator

// Example:
import "github.com/gomig/validator"
import "github.com/gomig/validator/translations"
v := NewValidator(t, "en")
translations.RegisterENValidationMessages(v) // EN
translations.RegisterFAValidationMessages(v)// FA (Persian)
```

## Usage

Validator interface contains following methods:

**Note:** You must pass app locale for generating localization message.

### Validator

Get original validator instance

```go
// Signature:
Validator() *validator.Validate
```

### AddValidation

Register new validator function.

```go
// Signature:
AddValidation(tag string, v validator.Func)
```

### AddTranslation

Register new translation message to validator translator.

```go
// Signature:
AddTranslation(locale string, key string, message string)
```

### Translate

Generate translation. used internally by validator!

```go
// Signature:
Translate(locale string, key string, placeholders map[string]string) string
```

### TranslateStruct

Generate translation for struct. used internally by validator!

```go
// Signature:
TranslateStruct(s any, locale string, key string, field string, placeholders map[string]string) string
```

### Struct

Validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.

**NOTE:** You can use `StructLocale` method for generate error message in non-default validator locale.

```go
// Signature:
Struct(s any) ErrorResponse
```

### StructExcept

Validates all fields except the ones passed in.

**NOTE:** You can use `StructExceptLocale` method for generate error message in non-default validator locale.

```go
// Signature:
StructExcept(s any, fields ...string) ErrorResponse
```

### StructPartial

Validates the fields passed in only, ignoring all others.

**NOTE:** You can use `StructPartialLocale` method for generate error message in non-default validator locale.

```go
// Signature:
StructPartial(s any, fields ...string) ErrorResponse
```

### Var

Validates a single variable using tag style validation. You must use ValidatorParam for customizing var validation.

**NOTE:** You can use `VarLocale` method for generate error message in non-default validator locale.

**Caution:** Name field is required and if not passed in params this validation generate a panic!

ValidatorParam fields:

- **Name:** this field name. e.g firstname
- **Title:** validation field title. works like vTitle tag
- **ParamTitle:** validation param title. works like vParam tag
- **Format:** format validation param if set to true. works like vFormat tag

```go
// Signature:
Var(params ValidatorParam, field any, tag string, messages map[string]string) ErrorResponse
```

**Note:** You can pass a message list to override default translation messages.

```go
v.Var(params, field, "validate:required", map[string]string{ "required":"enter your name" })
```

### VarWithValue

Validates a single variable, against another variable/field's value using tag style validation.

**NOTE:** You can use `VarWithValueLocale` method for generate error message in non-default validator locale.

```go
// Signature:
VarWithValue(params ValidatorParam, field any, other any, tag string, messages map[string]string) ErrorResponse
```

## Extra Validation Commands

Validator contains some extra validation commands you can register and use with your validator.

you can find this validations under `github.com/gomig/validator/validations` namespace.

### AlphaNum

Check if field is a string contains alpha and number. You can allow extra character by listing them in param.

```go
type Person struct{
    Field string `validate:"alnum=' -_'"`
}
```

### AlphaNumFa

Check if field is a string contains alpha (en and fa) and number (en only). You can allow extra character by listing them in param.

```go
type Person struct{
    Field string `validate:"alnumfa=' '"`
}
```

### Credit Card

Check if field is a credit card number.

```go
type Person struct{
    Field string `validate:"creditcard"`
}
```

### Identifier

Check if field is a valid numeric greater than 1.

```go
type Person struct{
    Field string `validate:"identifier"`
}
```

### ID Number

Check if field is a valid id number (number has 1-10 length).

```go
type Person struct{
    Field string `validate:"idnumber"`
}
```

### IP Port

Check if field is a valid ip:port string.

```go
type Person struct{
    Field string `validate:"ipport"`
}
```

### Jalaali

Check if field is a valid persian date string.

```go
type Person struct{
    Field string `validate:"jalaali"`
}
```

### Mobile

Check if field is a valid persian mobile number.

```go
type Person struct{
    Field string `validate:"mobile"`
}
```

### National Code

Check if field is a valid persian national code.

```go
type Person struct{
    Field string `validate:"nationalcode"`
}
```

### Postal Code

Check if field is a valid persian postal code.

```go
type Person struct{
    Field string `validate:"postalcode"`
}
```

### Tel

Check if field is a valid persian tel.

```go
type Person struct{
    Field string `validate:"tel"`
}
```

### Unsigned

Check if field is a valid unsigned number (0 or greater).

```go
type Person struct{
    Field string `validate:"unsigned"`
}
```

### Username

Check if field is a valid username (can contains 0-9, a-a, A-Z dash, dot and underscore).

```go
type Person struct{
    Field string `validate:"username"`
}
```

### UUID

Check if field is a valid uuid.

```go
type Person struct{
    Field string `validate:"uuid"`
}
```

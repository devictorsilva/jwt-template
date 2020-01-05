# JWT Generator 

Is a basic implementation of token generation API.

## Installation

We are using [go mod](https://blog.golang.org/using-go-modules) to manage all dependencies on this project. So you are free to run:

```bash
foo@bar:~$ go build -o ./cmd/jwt-generator --ldflags="-s -w" ./cmd/jwt-generator/
```

And the jwt-generator executable will be created on the `cmd/jwt-generator` folder.

The flags ( `--ldflags="-s -w"` ) is used to shrink the code, optinaly you can use the flags and [UPX](https://upx.github.io/) to get all power of shrinking.

## Usage

To use this program run the follow commands:

```bash
foo@bar:~$ cd ./cmd/jwt-generator/
foo@bar:~$ ./jwt-generator
```

Please make sure the `jwt-generator` is an executable.

### Endpoints

| HTTP METHOD | ENDPOINT       | DESCRIPTION                                    |
| ------ | -------------- | ---------------------------------------------- |
| GET    | `/hmac-token`  | Generate JWT token HMAC encoded (HS256,HS...)  |
| GET    | `/rsa-token`   | Generate JWT token RSA encoded (RS256,RS...)   |
| GET    | `/ecdsa-token` | Generate JWT token ECDSA encoded (ES256,ES...) |
| GET    | `/ping`        | For fun purposes and life check                |

### Request Structure

For the endpoints `/hmac-token`, `/rsa-token` and `/ecdsa-token` the request is:

#### Headers:

- **Content-Type** : `application/json`

#### Body:

The body is a json formated, with the following params:

- **method**: Method refers to method used to encode the claim. ( Required, string ) 
- **aud**: The "aud" ( audience ) claim identifies the recipients that the JWT is intended for. ( Optional, string )
- **iss**: The "iss" ( issuer ) claim identifies the principal that issued the JWT. ( Optional, string )
- **sub**: The "sub" ( subject ) claim identifies the principal that is the subject of the JWT. ( Optional, string )

- **nbf_in**: , nbf_in refes to the "nbf" (Not Before) claim that identifies the time before which the JWT MUST NOT be accepted for processing, if it's comes the calc to generate "nbf" is now + nbf_in (seconds). ( Optional, positive int )

- **exp_in**: exp_in refes to The "exp" (expiration time) claim identifies the expiration time on or after which the JWT MUST NOT be accepted for processing, if it's comes the calc to generate "exp" is now + exp_in (seconds). ( Optional, positive int )

----------------
The `method` depends of the encode.

| ENCODE | SUPORTED METHODS       |
| ------ | ---------------------- |
| HMAC   | HS256, HS384 and HS512 |
| RSA    | RS256, RS384 and RS512 |
| ECDSA  | ES256, ES384 and ES512 |

---------------------

For `/ping` endpoint is:

- No content

### Response Structure

For all endpoint the response structure is:

#### Headers:

- **Content-Type** : `application/json`

- **Authorizarion** : Bearer &lt;token> ( Maybe comes, if has no errors )

#### Body:

- **message** : Friendly message of the error / success ( Always comes )  

- **error** : Error message from the code ( Maybe comes, if has one error )

### Examples

#### Simple Request 

**HTTP METHOD**: 
- GET

**ENDPOINT**: 
- `/hmac-token`

**Headers**:

- **Content-Type** : `application/json`

**BODY**:
```json
{
    "method" : "HS256"
}
```

#### Success Response

**HTTP Status**: 
- 200 OK

**Headers**:

- **Content-Type** : `application/json`
- **Authorization** : `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzgxNzAzNDYsImp0aSI6IjUyYTg1YTIzLWY4N2EtNGU2MS1iZDgwLTJmMjQxYWVmOWE3MSIsImlhdCI6MTU3ODE3MDA0Nn0.Z4YHvWnIwUvMres70iwLhv9decVSrVD7YMlInP7PtZA`

**BODY**:
```json
{
    "message": "token returned, see the Authorizarion header",
    "error": null
}
```

**Decoded Token Payload**:
```json
{
  "exp": 1578170346,
  "jti": "52a85a23-f87a-4e61-bd80-2f241aef9a71",
  "iat": 1578170046
}
```

--------------

#### Bad Request 

**HTTP METHOD**: 
- GET

**ENDPOINT**: 
- `/rsa-token`

**Headers**:

- **Content-Type** : `application/json`

**BODY**:
```json
{
    "method" : "HS256"
}
```

#### Error Response

**HTTP Status**: 
- 400 Bad Request

**Headers**:

- **Content-Type** : `application/json`

**BODY**:
```json
{
    "error": "invalid encode method",
    "message": "please use a valid RSA encode method"
}
```

--------------

#### Request With Optional Params

**HTTP METHOD**: 
- GET

**ENDPOINT**: 
- `/ecdsa-token`

**Headers**:

- **Content-Type** : `application/json`

**BODY**:
```json
{
    "method" : "ES256",
    "nbf_in": 600,
    "exp_in": 60,
    "iss": "jwt-generator",
    "sub": "devictorsilva",
    "aud": "jwt-validator"
}
```

#### Success Response

**HTTP Status**: 
- 200 OK

**Headers**:

- **Content-Type** : `application/json`
- **Authorization** : `Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJqd3QtdmFsaWRhdG9yIiwiZXhwIjoxNTc4MTc0OTk3LCJqdGkiOiJjN2VmNTU2Ni1lYjRjLTRhNDUtOWQ5YS0yNjA5NjE3NWFjZTUiLCJpYXQiOjE1NzgxNzQ5MzcsImlzcyI6Imp3dC1nZW5lcmF0b3IiLCJuYmYiOjE1NzgxNzU1MzcsInN1YiI6ImRldmljdG9yc2lsdmEifQ.ygC2urSPdSiNZP_rYRHScpGluzOYxzm_NWzpui28XMuDrcHc_6atRhu0-qzp1o5JL0h9U2HtW9IxNKReE6S8iw`

**BODY**:
```json
{
    "message": "token returned, see the Authorizarion header",
    "error": null
}
```

**Decoded Token Payload**:
```json
{
  "aud": "jwt-validator",
  "exp": 1578174997,
  "jti": "c7ef5566-eb4c-4a45-9d9a-26096175ace5",
  "iat": 1578174937,
  "iss": "jwt-generator",
  "nbf": 1578175537,
  "sub": "devictorsilva"
}
```

-------------------


#### Ping Request

**HTTP METHOD**: 
- POST

**ENDPOINT**: 
- `/ping`

**Headers**:

- No Content

**BODY**:

- No Content

#### Success Response

**HTTP Status**: 
- 200 OK

**Headers**:

- **Content-Type** : `application/json`

**BODY**:
```json
{
    "message": "pong!",
    "error": null
}
```
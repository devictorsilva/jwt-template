# JWT Validator 

Is a basic implementation of token validation API.

## Installation

We are using [go mod](https://blog.golang.org/using-go-modules) to manage all dependencies on this project. So you are free to run:

```bash
foo@bar:~$ go build -o ./cmd/jwt-validator --ldflags="-s -w" ./cmd/jwt-validator/
```

And the jwt-validator executable will be created on the `cmd/jwt-validator` folder.

The flags ( `--ldflags="-s -w"` ) is used to shrink the code, optinaly you can use the flags and [UPX](https://upx.github.io/) to get all power of shrinking.

## Usage

To use this program run the follow commands:

```bash
foo@bar:~$ cd ./cmd/jwt-validator/
foo@bar:~$ ./jwt-validator
```

Please make sure the `jwt-validator` is an executable.

### Endpoints

| HTTP METHOD | ENDPOINT        | DESCRIPTION                                         |
| ----------- | --------------- | --------------------------------------------------- |
| POST        | `/ping`         | For fun purposes and life check                     |
| POST        | `/private/ping` | Validate the token and return a authenticated pong! |

### Request Structure

For `/ping` endpoint is:

- No content

For `/private/ping` endpoint is:

**Headers**:

- **Authorizarion** : Bearer &lt;token> 

### Response Structure

For all endpoint the response structure is:

#### Headers:

- **Content-Type** : `application/json`

#### Body:

- **message** : Friendly message of the error / success ( Always comes )  

- **error** : Error message from the code ( Maybe comes, if has one error )

### Examples

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

---------------


#### Authenticated Ping Request

**HTTP METHOD**: 
- POST

**ENDPOINT**: 
- `/private/ping`

**Headers**:

- **Authorization** : `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzgxOTUwMDAsImp0aSI6IjUyYTg1YTIzLWY4N2EtNGU2MS1iZDgwLTJmMjQxYWVmOWE3MSIsImlhdCI6MTU3ODE5MDk4Mn0.l_kBg63EruPl1nSGw-iAR4zwVeRiUVX-zVHNaPQdzGc`

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
    "message": "authenticated pong!",
    "error": null
}
```
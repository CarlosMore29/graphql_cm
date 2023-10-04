# Ayuda

Consulta una API Graphql (queries y mutaciones) con ayuda de la siguiente siguiente liberia github.com/machinebox/graphql

# Uso

Para inicializar un objeto

```go
var objAPIQL *graphql_cm.APIQL = graphql_cm.InitAPIQL()
```

Por defaulf usará estas variables de entorno

```go
os.Getenv("APIQL_URL")
os.Getenv("APIQL_SECRET")
```

Si necesitas cambiar las credenciales (URL & TOKEN) USA

```go
SetCredentials
```

## v1.0.0

### QueryRequest

Diseñada para ejecutar queries, manda el string query y una array de inputs (InputVariable),

```go
objAPIQL.QueryRequest(query string, input []InputVariable)
```

Estructura creada para enviar el valor esperado (string o int)

```go
type InputVariable struct {
    Key         string
    ValueString string
    ValueInt    int
    ValueBool   bool
}
```

### MutationRequest

Diseñada para ejecutar mutaciones, lo que envies en variables va dentro de un obj "input"

```go
objAPIQL.MutationRequest(query string, variables interface{})
```

## v1.1.0

### QueryRequest

Se cambio la forma en la que enviamos datos atraves de una query

```go
type InputVariable struct {
    Key   string
    Value Interface{}
}
```
### MutationRequest

funcion que ejecuta mutaciones ahora acepta que cambies el nombre a la raiz del objeto input

```go
objAPIQL.MutationRequest(query, inputRootName string, variables interface{})
```
objeto input
```json
{
  <inputRootName>: {
    <variables>
    }
}
```

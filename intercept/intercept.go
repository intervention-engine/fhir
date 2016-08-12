package intercept

import (
    "fmt"
    "reflect"
)

func PostInterceptor(resource interface{}) {
    fmt.Printf("POST intercepted: %s\n", getResourceType(resource))
}

func PutInterceptor(resource interface{}) {
    fmt.Printf("PUT intercepted: %s\n", getResourceType(resource))
}

func DeleteInterceptor(resource interface{}) {
    fmt.Printf("DELETE intercepted: %s\n", getResourceType(resource))
}

func getResourceType(resource interface{}) string {
    resType := reflect.TypeOf(resource).Elem().Name()
    return resType
}

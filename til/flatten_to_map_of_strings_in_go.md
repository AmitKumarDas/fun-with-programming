#### References
```yaml
- https://github.com/gavinbunney/terraform-provider-kubectl/blob/master/flatten/flatten.go
- https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go
```

```go
// Flatten takes a structure and turns into a flat map[string]string.
//
// Based on https://github.com/hashicorp/terraform/blob/master/flatmap/flatten.go
//
// Structs are not supported. Therefore, it can only be slices, maps, primitives,
// and any combination of those together.
func Flatten(thing map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for k, raw := range thing {
		if raw == nil || k == ""{
			continue
		}

		flatten(result, k, reflect.ValueOf(raw))
	}
	return result
}

// Note: key is understood as a prefix due to flattening
func flatten(result map[string]string, prefix string, v reflect.Value) {
	if v.Kind() == reflect.Invalid {
		return
	}

	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			result[prefix] = "true"
		} else {
			result[prefix] = "false"
		}
	case reflect.Int:
		result[prefix] = fmt.Sprintf("%d", v.Int())
	case reflect.Map:
		flattenMap(result, prefix, v)
	case reflect.Slice:
		flattenSlice(result, prefix, v)
	case reflect.String:
		result[prefix] = v.String()
	default:
		result[prefix] = fmt.Sprintf("%v", v) // structs, Int64, Float64, etc.
	}
}

func flattenMap(result map[string]string, prefix string, v reflect.Value) {
	for _, k := range v.MapKeys() {
		if k.Kind() == reflect.Interface {
			k = k.Elem()
		}

		key := fmt.Sprintf("%v", k)
		flatten(result, fmt.Sprintf("%s.%s", prefix, key), v.MapIndex(k)) // TIL
	}
}

func flattenSlice(result map[string]string, prefix string, v reflect.Value) {
	prefix = prefix + "."

	result[prefix+"#"] = fmt.Sprintf("%d", v.Len()) // Nice !!!
	for i := 0; i < v.Len(); i++ {
		flatten(result, fmt.Sprintf("%s%d", prefix, i), v.Index(i)) // TIL
	}
}

func TestFlattenMap(t *testing.T) {
	testCases := []struct {
		description string
		test        map[string]interface{}
		expected    map[string]string
	}{
		{
			description: "One level map",
			test: map[string]interface{}{
				"atest": "test",
				"meta": map[string]interface{}{
					"annotations": map[string]string{
						"helm.sh/hook": "crd-install",
					},
				},
			},
			expected: map[string]string{
				"atest":                         "test",
				"meta.annotations.helm.sh/hook": "crd-install",
			},
		},
		{
			description: "One level slice",
			test: map[string]interface{}{
				"my-slice": []string{"first", "second"},
			},
			expected: map[string]string{
				"my-slice.#": "2",
				"my-slice.0": "first",
				"my-slice.1": "second",
			},
		},
		{
			description: "Map with slice elements",
			test: map[string]interface{}{
				"meta": map[string]interface{}{
					"my-slice": []string{"first", "second"},
				},
			},
			expected: map[string]string{
				"meta.my-slice.#": "2",
				"meta.my-slice.0": "first",
				"meta.my-slice.1": "second",
			},
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.description, func(t *testing.T) {
			result := Flatten(tcase.test)
			assert.Equal(t, tcase.expected, result)
		})
	}
}
```

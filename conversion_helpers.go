package main

func interfaceSliceToStringSlice(in []interface{}) []string {
	out := make([]string, len(in))
	for i := range in {
		out[i] = in[i].(string)
	}
	return out
}

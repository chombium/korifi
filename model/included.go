package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PaesslerAG/jsonpath"
)

type IncludedResource struct {
	Type     string
	Resource any
}

func (r IncludedResource) SelectJSONPaths(paths ...string) (IncludedResource, error) {
	resourceBytes, err := json.Marshal(r.Resource)
	if err != nil {
		return IncludedResource{}, fmt.Errorf("failed to marshal resource: %w", err)
	}

	resourceMap := map[string]any{}
	if err = json.Unmarshal(resourceBytes, &resourceMap); err != nil {
		return IncludedResource{}, fmt.Errorf("failed to unmarshal resource: %w", err)
	}

	if len(paths) == 0 {
		return IncludedResource{
			Type:     r.Type,
			Resource: resourceMap,
		}, nil
	}

	includedResource := map[string]any{}
	for _, path := range paths {
		value, err := jsonpath.Get("$."+path, resourceMap)
		if err != nil {
			return IncludedResource{}, fmt.Errorf("failed to select %q from %q: %w", path, resourceMap, err)
		}

		pathElements := strings.Split(path, ".")
		includedResource[pathElements[0]] = buildTree(pathElements[1:], value)
	}

	return IncludedResource{
		Type:     r.Type,
		Resource: includedResource,
	}, nil
}

func buildTree(pathElements []string, leafValue any) any {
	if len(pathElements) == 0 {
		return leafValue
	}

	return map[string]any{
		pathElements[0]: buildTree(pathElements[1:], leafValue),
	}
}
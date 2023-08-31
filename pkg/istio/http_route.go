package istio

import (
	"github.com/getkin/kin-openapi/openapi3"
	istioapi "istio.io/api/networking/v1beta1"
)

func HTTPRoutesFromOpenAPI(oasDoc *openapi3.T, destination *istioapi.Destination, pathMatchPrefix bool, pathMatchRegex bool) ([]*istioapi.HTTPRoute, error) {
	httpRoutes := []*istioapi.HTTPRoute{}

	// Path based routing
	for path, pathItem := range oasDoc.Paths {

		var pathMatchType *istioapi.StringMatch

		if pathMatchRegex {
			pathMatchType = &istioapi.StringMatch{
				MatchType: &istioapi.StringMatch_Regex{Regex: path},
			}
		} else if pathMatchPrefix {
			pathMatchType = &istioapi.StringMatch{
				MatchType: &istioapi.StringMatch_Prefix{Prefix: path},
			}
		} else {
			pathMatchType = &istioapi.StringMatch{
				MatchType: &istioapi.StringMatch_Exact{Exact: path},
			}
		}

		for opVerb, operation := range pathItem.Operations() {

			queryParams := make(map[string]*istioapi.StringMatch)
			//fmt.Println(json.Marshal(pathItem.Parameters))
			for _, opParam := range operation.Parameters {
				queryParams[opParam.Value.Name] = &istioapi.StringMatch{MatchType: &istioapi.StringMatch_Regex{Regex: ".*"}}
			}

			httpRoute := &istioapi.HTTPRoute{
				// TODO(eastizle): OperationID can be null, fallback to some custom name
				Name: operation.OperationID,
				Match: []*istioapi.HTTPMatchRequest{
					{
						Uri: pathMatchType,
						Method: &istioapi.StringMatch{
							MatchType: &istioapi.StringMatch_Exact{Exact: opVerb},
						},
						QueryParams: queryParams,
					},
				},
				Route: []*istioapi.HTTPRouteDestination{{Destination: destination}},
			}
			httpRoutes = append(httpRoutes, httpRoute)
		}
	}

	return httpRoutes, nil
}

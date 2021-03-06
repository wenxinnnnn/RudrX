package traitdefinition

var SimpleRollout = `apiVersion: core.oam.dev/v1alpha2
kind: TraitDefinition
metadata:
  name: simplerollouttraits.extend.oam.dev
  annotations:
    "definition.oam.dev/apiVersion": "extend.oam.dev/v1alpha2"
    "definition.oam.dev/kind": "SimpleRolloutTrait"
spec:
  revisionEnabled: true
  appliesToWorkloads:
    - core.oam.dev/v1alpha2.ContainerizedWorkload
    - deployments.apps
  definitionRef:
    name: simplerollouttraits.extend.oam.dev
  extension:
    template: |
      #Template: {
      	apiVersion: "extend.oam.dev/v1alpha2"
      	kind:       "SimpleRolloutTrait"
      	spec: {
      		replica:        rollout.replica
      		maxUnavailable: rollout.maxUnavailable
      		batch:          rollout.batch
      	}
      }
      rollout: {
      	replica:        *3 | int
      	maxUnavailable: *1 | int
      	batch:          *2 | int
      }
`

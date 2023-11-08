# operator-sdk-defect

Sample operator to demo an issue with `operator-sdk generate kustomize manifests` which forget some fields in CSV when the API is split in multiple Go packages.

See how the CSV in config/manifests/bases/operator-sdk-defect.clusterserviceversion.yaml doesn't include the `options.hidden1` property and its `com.tectonic.ui:hidden` flag.

(Defect pending)

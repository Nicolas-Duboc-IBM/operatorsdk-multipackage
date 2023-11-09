# operator-sdk-defect

Sample operator to demo an issue with `operator-sdk generate kustomize manifests` which forgets some fields in CSV when the API is splitted in multiple Go packages.

See how the CSV in config/manifests/bases/operator-sdk-defect.clusterserviceversion.yaml doesn't include the `options.hidden1` property and its `com.tectonic.ui:hidden` flag.

Issue reported as https://github.com/operator-framework/operator-sdk/issues/6624

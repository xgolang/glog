/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// test file for unstructured logging static check tool unit tests.

package testdata

import (
	klog "k8s.io/klog/v2"
)

func printUnstructuredLog() {
	klog.V(1).Infof("test log")      // want `unstructured logging function "Infof" should not be used`
	klog.Infof("test log")           // want `unstructured logging function "Infof" should not be used`
	klog.Info("test log")            // want `unstructured logging function "Info" should not be used`
	klog.Infoln("test log")          // want `unstructured logging function "Infoln" should not be used`
	klog.InfoDepth(1, "test log")    // want `unstructured logging function "InfoDepth" should not be used`
	klog.Warning("test log")         // want `unstructured logging function "Warning" should not be used`
	klog.Warningf("test log")        // want `unstructured logging function "Warningf" should not be used`
	klog.WarningDepth(1, "test log") // want `unstructured logging function "WarningDepth" should not be used`
	klog.Error("test log")           // want `unstructured logging function "Error" should not be used`
	klog.Errorf("test log")          // want `unstructured logging function "Errorf" should not be used`
	klog.Errorln("test log")         // want `unstructured logging function "Errorln" should not be used`
	klog.ErrorDepth(1, "test log")   // want `unstructured logging function "ErrorDepth" should not be used`
	klog.Fatal("test log")           // want `unstructured logging function "Fatal" should not be used`
	klog.Fatalf("test log")          // want `unstructured logging function "Fatalf" should not be used`
	klog.Fatalln("test log")         // want `unstructured logging function "Fatalln" should not be used`
	klog.FatalDepth(1, "test log")   // want `unstructured logging function "FatalDepth" should not be used`

}

func printStructuredLog() {
	klog.InfoS("test log")
	klog.ErrorS(nil, "test log")
	klog.InfoS("Starting container in a pod", "containerID", "containerID", "pod")                // want `Additional arguments to InfoS should always be Key Value pairs. Please check if there is any key or value missing.`
	klog.ErrorS(nil, "Starting container in a pod", "containerID", "containerID", "pod")          // want `Additional arguments to ErrorS should always be Key Value pairs. Please check if there is any key or value missing.`
	klog.InfoS("Starting container in a pod", "测试", "containerID")                                // want `Key positional arguments "测试" are expected to be lowerCamelCase alphanumeric strings. Please remove any non-Latin characters.`
	klog.ErrorS(nil, "Starting container in a pod", "测试", "containerID")                          // want `Key positional arguments "测试" are expected to be lowerCamelCase alphanumeric strings. Please remove any non-Latin characters.`
	klog.InfoS("Starting container in a pod", 7, "containerID")                                   // want `Key positional arguments are expected to be inlined constant strings. Please replace 7 provided with string value`
	klog.ErrorS(nil, "Starting container in a pod", 7, "containerID")                             // want `Key positional arguments are expected to be inlined constant strings. Please replace 7 provided with string value`
	klog.InfoS("Starting container in a pod", map[string]string{"test1": "value"}, "containerID") // want `Key positional arguments are expected to be inlined constant strings. `
	testKey := "a"
	klog.ErrorS(nil, "Starting container in a pod", testKey, "containerID") // want `Key positional arguments are expected to be inlined constant strings. `
	klog.InfoS("test: %s", "testname")                                      // want `structured logging function "InfoS" should not use format specifier "%s"`
	klog.ErrorS(nil, "test no.: %d", 1)                                     // want `structured logging function "ErrorS" should not use format specifier "%d"`
}

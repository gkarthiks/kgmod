/*
Package utils holds the utility files and methods for all the commands
and subcommands og kgmod
*/
/*
Copyright © 2020 Karthikeyan Govindaraj <github.gkarthiks@gmail.com>

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
package utils

// GoExecPath holds the go executable path
var (
	GoExecPath string
	DockerPath string
	HelmPath   string
)

// BuildVersion holds the version of the kgmod
var BuildVersion = "development"

// ConfigFileLocation basic config location
var ConfigFileLocation = "https://raw.githubusercontent.com/gkarthiks/k8s-dumps/master/kgmod.yaml"

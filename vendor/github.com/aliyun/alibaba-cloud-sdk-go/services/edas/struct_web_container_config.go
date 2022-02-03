package edas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// WebContainerConfig is a nested struct in edas response
type WebContainerConfig struct {
	ContextInputType     string `json:"ContextInputType" xml:"ContextInputType"`
	ContextPath          string `json:"ContextPath" xml:"ContextPath"`
	HttpPort             int    `json:"HttpPort" xml:"HttpPort"`
	MaxThreads           int    `json:"MaxThreads" xml:"MaxThreads"`
	ServerXml            string `json:"ServerXml" xml:"ServerXml"`
	UriEncoding          string `json:"UriEncoding" xml:"UriEncoding"`
	UseAdvancedServerXml bool   `json:"UseAdvancedServerXml" xml:"UseAdvancedServerXml"`
	UseBodyEncoding      bool   `json:"UseBodyEncoding" xml:"UseBodyEncoding"`
	UseDefaultConfig     bool   `json:"UseDefaultConfig" xml:"UseDefaultConfig"`
}
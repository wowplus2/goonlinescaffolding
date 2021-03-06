// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gos_serveconnbyte

const (
	HTML_tableheader = `
<tr>
<th>ConnData</th>
<th>APIstats</th>
<th>NotiStats</th>
<th>ErrorStats</th>
<th>Auth Cmds</th>
</tr>`

	HTML_row = `
<tr>
<td>{{$v.GetConnData}}</td>
<td>{{$v.GetAPIStat}}</td>
<td>{{$v.GetNotiStat}}</td>
<td>{{$v.GetErrorStat}}</td>
<td>{{$v.GetAuthorCmdList}}</td>
</tr>
`
)

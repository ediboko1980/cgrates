/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package config

import (
	"github.com/cgrates/cgrates/utils"
)

// HTTPCfg is the HTTP config section
type HTTPCfg struct {
	HTTPJsonRPCURL          string            // JSON RPC relative URL ("" to disable)
	DispatchersRegistrarURL string            // dispatcherH registrar service relative URL
	HTTPWSURL               string            // WebSocket relative URL ("" to disable)
	HTTPFreeswitchCDRsURL   string            // Freeswitch CDRS relative URL ("" to disable)
	HTTPCDRsURL             string            // CDRS relative URL ("" to disable)
	HTTPUseBasicAuth        bool              // Use basic auth for HTTP API
	HTTPAuthUsers           map[string]string // Basic auth user:password map (base64 passwords)
	ClientOpts              map[string]interface{}
}

// loadFromJsonCfg loads Database config from JsonCfg
func (httpcfg *HTTPCfg) loadFromJsonCfg(jsnHttpCfg *HTTPJsonCfg) (err error) {
	if jsnHttpCfg == nil {
		return nil
	}
	if jsnHttpCfg.Json_rpc_url != nil {
		httpcfg.HTTPJsonRPCURL = *jsnHttpCfg.Json_rpc_url
	}
	if jsnHttpCfg.Dispatchers_registrar_url != nil {
		httpcfg.DispatchersRegistrarURL = *jsnHttpCfg.Dispatchers_registrar_url
	}
	if jsnHttpCfg.Ws_url != nil {
		httpcfg.HTTPWSURL = *jsnHttpCfg.Ws_url
	}
	if jsnHttpCfg.Freeswitch_cdrs_url != nil {
		httpcfg.HTTPFreeswitchCDRsURL = *jsnHttpCfg.Freeswitch_cdrs_url
	}
	if jsnHttpCfg.Http_Cdrs != nil {
		httpcfg.HTTPCDRsURL = *jsnHttpCfg.Http_Cdrs
	}
	if jsnHttpCfg.Use_basic_auth != nil {
		httpcfg.HTTPUseBasicAuth = *jsnHttpCfg.Use_basic_auth
	}
	if jsnHttpCfg.Auth_users != nil {
		httpcfg.HTTPAuthUsers = *jsnHttpCfg.Auth_users
	}
	if jsnHttpCfg.Client_opts != nil {
		for k, v := range jsnHttpCfg.Client_opts {
			httpcfg.ClientOpts[k] = v
		}
	}
	return nil
}

func (httpcfg *HTTPCfg) AsMapInterface() (initialMP map[string]interface{}) {
	initialMP = map[string]interface{}{
		utils.HTTPJsonRPCURLCfg:          httpcfg.HTTPJsonRPCURL,
		utils.DispatchersRegistrarURLCfg: httpcfg.DispatchersRegistrarURL,
		utils.HTTPWSURLCfg:               httpcfg.HTTPWSURL,
		utils.HTTPFreeswitchCDRsURLCfg:   httpcfg.HTTPFreeswitchCDRsURL,
		utils.HTTPCDRsURLCfg:             httpcfg.HTTPCDRsURL,
		utils.HTTPUseBasicAuthCfg:        httpcfg.HTTPUseBasicAuth,
		utils.HTTPAuthUsersCfg:           httpcfg.HTTPAuthUsers,
		utils.HTTPClientOptsCfg:          httpcfg.ClientOpts,
	}
	return
}

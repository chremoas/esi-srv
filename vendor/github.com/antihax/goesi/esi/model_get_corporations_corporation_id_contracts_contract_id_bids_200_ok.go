/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.9
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

/* A list of GetCorporationsCorporationIdContractsContractIdBids200Ok. */
//easyjson:json
type GetCorporationsCorporationIdContractsContractIdBids200OkList []GetCorporationsCorporationIdContractsContractIdBids200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdContractsContractIdBids200Ok struct {
	Amount   float32   `json:"amount,omitempty"`    /* The amount bid, in ISK */
	BidId    int32     `json:"bid_id,omitempty"`    /* Unique ID for the bid */
	BidderId int32     `json:"bidder_id,omitempty"` /* Character ID of the bidder */
	DateBid  time.Time `json:"date_bid,omitempty"`  /* Datetime when the bid was placed */
}

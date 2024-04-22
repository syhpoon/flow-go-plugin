/*
 * Flow Go SDK
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package flow

import "time"

// Block is a set of state mutations applied to the Flow blockchain.
type Block struct {
	BlockHeader
	BlockPayload
}

// BlockHeader is a summary of a full block.
type BlockHeader struct {
	ID        Identifier
	ParentID  Identifier
	Height    uint64
	Timestamp time.Time
	Status    BlockStatus
}

// BlockStatus represents the status of a block.
type BlockStatus int

const (
	// BlockStatusUnknown indicates that the block status is not known.
	BlockStatusUnknown BlockStatus = iota
	// BlockStatusFinalized is the status of a finalized block.
	BlockStatusFinalized
	// BlockStatusSealed is the status of a sealed block.
	BlockStatusSealed
)

func BlockStatusFromString(s string) BlockStatus {
	switch s {
	case "BLOCK_FINALIZED":
		return BlockStatusFinalized
	case "BLOCK_SEALED":
		return BlockStatusSealed
	default:
		return BlockStatusUnknown
	}
}

// BlockPayload is the full contents of a block.
//
// A payload contains the collection guarantees and seals for a block.
type BlockPayload struct {
	CollectionGuarantees []*CollectionGuarantee
	Seals                []*BlockSeal
}

// BlockSeal is the attestation by verification nodes that the transactions in a previously
// executed block have been verified.
type BlockSeal struct {
	// The ID of the block this Seal refers to (which will be of lower height than this block)
	BlockID Identifier

	// The ID of the execution receipt generated by the Verifier nodes; the work of verifying a
	// block produces the same receipt among all verifying nodes
	ExecutionReceiptID Identifier
}

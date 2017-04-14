/*

Copyright 2017 Continusec Pty Ltd

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

package api

import (
	"log"

	"github.com/continusec/verifiabledatastructures/pb"
	"github.com/golang/protobuf/proto"
)

type LocalService struct {
	Mutator      MutatorService // if nil, we send to ourself which simply applies instantly, blocking until return
	AccessPolicy AuthorizationOracle
	Reader       StorageReader
}

func ApplyMutation(db KeyWriter, mut *pb.Mutation) error {
	log.Printf("Mutation: %s\n", proto.CompactTextString(mut))
	switch {
	case mut.LogAddEntry != nil:
		return applyLogAddEntry(db, mut.LogAddEntry)
	default:
		return ErrNotImplemented
	}
}

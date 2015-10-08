/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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

package inst

import (
	"github.com/outbrain/golib/log"
)

type PostponedFunctionsContainer struct {
	PostponedFunctions [](func() error)
}

func NewPostponedFunctionsContainer() *PostponedFunctionsContainer {
	postponedFunctionsContainer := &PostponedFunctionsContainer{}
	postponedFunctionsContainer.PostponedFunctions = [](func() error){}
	return postponedFunctionsContainer
}

func (this *PostponedFunctionsContainer) AddPostponedFunction(f func() error) {
	this.PostponedFunctions = append(this.PostponedFunctions, f)
}

func (this *PostponedFunctionsContainer) InvokePostponed() (err error) {
	if len(this.PostponedFunctions) == 0 {
		return
	}
	log.Debugf("PostponedFunctionsContainer: invoking %+v postponed functions", len(this.PostponedFunctions))
	for _, postponedFunction := range this.PostponedFunctions {
		ferr := postponedFunction()
		if err == nil {
			err = ferr
		}
	}
	return err
}

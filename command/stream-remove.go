package command

import (
	"fmt"
	"log"
	"lsf"
	"lsf/schema"
	"lsf/system"
)

const removeStreamCmdCode lsf.CommandCode = "stream-remove"

type removeStreamOptionsSpec struct {
	global BoolOptionSpec
	id     StringOptionSpec
}

var removeStream *lsf.Command
var removeStreamOptions *removeStreamOptionsSpec

func init() {

	removeStream = &lsf.Command{
		Name:  removeStreamCmdCode,
		About: "Remove a new log stream",
		Init:  verifyRemoveStreamRequiredOpts,
		Run:   runRemoveStream,
		Flag:  FlagSet(removeStreamCmdCode),
	}
	removeStreamOptions = &removeStreamOptionsSpec{
		global: NewBoolFlag(removeStream.Flag, "g", "gg", false, "ggg", false),
		id:     NewStringFlag(removeStream.Flag, "s", "stream-id", "", "unique identifier for stream", true),
	}
}
func verifyRemoveStreamRequiredOpts(env *lsf.Environment, args ...string) error {
	if e := verifyRequiredOption(removeStreamOptions.id); e != nil {
		return e
	}
	return nil
}

func runRemoveStream(env *lsf.Environment, args ...string) error {

	id := schema.StreamId(*removeStreamOptions.id.value)

	// check if exists
	docid := system.DocId(fmt.Sprintf("stream.%s.stream", id))
	doc, e := env.LoadDocument(docid)
	if e != nil || doc == nil {
		return lsf.E_NOTEXISTING
	}

	// lock lsf port's "streams" resource
	lockid := env.ResourceId("streams")
	//	log.Printf("DEBUG: runAddStream: lockid: %q", lockid)
	lock, ok, e := system.LockResource(lockid, "add stream "+string(id))
	if e != nil {
		return e
	}
	if !ok {
		return fmt.Errorf("error - could not lock resource %q for stream add op", string(id))
	}
	defer lock.Unlock()

	panic("command.runRemoveStream() not impelemented - TODO system.Registrar.DeleteDocument(id)")
}

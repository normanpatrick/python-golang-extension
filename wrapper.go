/*
*   Simple golang python extension implementation.
*
*   Copyright (c) 2017, Norman Patrick
*
*   Permission to use, copy, modify, and distribute this software for any
*   purpose with or without fee is hereby granted, provided that the above
*   copyright notice and this permission notice appear in all copies.
*
*   THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
*   WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
*   MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
*   ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
*   WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
*   ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
*   OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

// #cgo pkg-config: python3
// #include <Python.h>
// int PyArg_ParseTuple_List(PyObject *, PyObject **);
// PyObject* Py_ReturnErrorMessage(char*);
// int PyArg_ParseTuple_StrAndList(PyObject *, char**, PyObject**);
import "C"

//export experiment
func experiment(self, args *C.PyObject) *C.PyObject {
	var listObj *C.PyObject
	if C.PyArg_ParseTuple_List(args, &listObj) == 0 {
		return nil
	}
	itemCount := int(C.PyList_Size(listObj))
	var flist []string
	for i := 0; i < itemCount; i++ {
		compTuple := C.PyList_GetItem(listObj, C.Py_ssize_t(i))
		var objectCompName *C.PyObject
		objectCompName = C.PySequence_GetItem(compTuple, 0)
		if objectCompName == nil {
			return C.Py_ReturnErrorMessage(C.CString("Failed to get list items"))
		}
		pStrObj := C.PyUnicode_AsUTF8String(compTuple)
		item := C.GoString(C.PyBytes_AsString(pStrObj))
		flist = append(flist, item)
	}
	res := myExperimentFn1(flist)
	retList := C.PyList_New(C.Py_ssize_t(len(res)))
	for i, iValue := range res {
		C.PyList_SetItem(retList, C.Py_ssize_t(i),
			C.PyLong_FromLong(C.long(iValue)))
	}
	return retList
}

//export experiment2
func experiment2(self, args *C.PyObject) *C.PyObject {
	var listObj *C.PyObject
	var bName *C.char
	if C.PyArg_ParseTuple_StrAndList(args, &bName, &listObj) == 0 {
		return nil
	}

	pyObjToGoString := func(pobj *C.PyObject) string {
		pStrObj_ := C.PyUnicode_AsUTF8String(pobj)
		item_ := C.GoString(C.PyBytes_AsString(pStrObj_))
		return item_
	}

	itemCount := int(C.PyList_Size(listObj))
	var flist []string
	for i := 0; i < itemCount; i++ {
		compTuple := C.PyList_GetItem(listObj, C.Py_ssize_t(i))
		var objectCompName *C.PyObject
		objectCompName = C.PySequence_GetItem(compTuple, 0)
		if objectCompName == nil {
			return C.Py_ReturnErrorMessage(C.CString("Failed to get list items"))
		}
		flist = append(flist, pyObjToGoString(compTuple))
	}
	res := myExperimentFn2(C.GoString(bName), flist)
	retList := C.PyList_New(C.Py_ssize_t(len(res)))
	for i, iValue := range res {
		C.PyList_SetItem(retList, C.Py_ssize_t(i),
			C.PyLong_FromLong(C.long(iValue)))
	}
	return retList
}

func main() {}

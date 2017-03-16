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

#include <Python.h>

PyObject * experiment(PyObject *, PyObject *);
PyObject * experiment2(PyObject *, PyObject *);

int PyArg_ParseTuple_List(PyObject * args, PyObject** pyList) {  
    return PyArg_ParseTuple(args, "O!", &PyList_Type, pyList);
}

int PyArg_ParseTuple_StrAndList(PyObject * args, char** pyStr,
                                PyObject** pyList) {  
    return PyArg_ParseTuple(args, "sO!", pyStr, &PyList_Type, pyList);
}

PyObject* Py_ReturnErrorMessage(char* msg) {
    return Py_BuildValue("s", msg);
}

static PyMethodDef Methods[] = {  
    {"experiment", experiment, METH_VARARGS,
     "Experiments....take a str list, return a number list"},
    {"experiment2", experiment2, METH_VARARGS,
     "Experiments....take a str and a str list, return a number list"},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef experimentmodule = {  
   PyModuleDef_HEAD_INIT, "experiment", NULL, -1, Methods
};

PyMODINIT_FUNC  
PyInit_experiment(void)  
{
    return PyModule_Create(&experimentmodule);
}

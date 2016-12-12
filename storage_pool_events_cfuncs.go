package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <assert.h>
#include "storage_pool_compat.h"
#include "storage_pool_events_cfuncs.h"
#include "callbacks_cfuncs.h"
#include <stdint.h>

extern void storagePoolEventLifecycleCallback(virConnectPtr, virStoragePoolPtr, int, int, int);
void storagePoolEventLifecycleCallback_cgo(virConnectPtr c, virStoragePoolPtr d,
                                           int event, int detail, void *data)
{
    storagePoolEventLifecycleCallback(c, d, event, detail, (int)(intptr_t)data);
}

int virConnectStoragePoolEventRegisterAny_cgo(virConnectPtr c,  virStoragePoolPtr d,
                                              int eventID, virConnectStoragePoolEventGenericCallback cb,
                                              long goCallbackId) {
#if LIBVIR_VERSION_NUMBER < 2000000
    assert(0); // Caller should have checked version
#else
    void* id = (void*)goCallbackId;
    return virConnectStoragePoolEventRegisterAny(c, d, eventID, cb, id, freeGoCallback_cgo);
#endif
}

*/
import "C"
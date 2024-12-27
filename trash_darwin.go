//go:build darwin
// +build darwin

package trash

/*
#cgo LDFLAGS: -framework Foundation
#include <Foundation/Foundation.h>

void moveToTrash(const char* path) {
    @autoreleasepool {
        NSString *nsPath = [NSString stringWithUTF8String:path];
        NSURL *url = [NSURL fileURLWithPath:nsPath];
        NSError *error = nil;
        [[NSFileManager defaultManager] trashItemAtURL:url resultingItemURL:nil error:&error];
        if (error) {
            NSLog(@"Error: %@", error);
        }
    }
}
*/
import "C"
import "unsafe"

func MoveToTrash(filePath string) error {
	cPath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cPath))
	C.moveToTrash(cPath)
	return nil
}

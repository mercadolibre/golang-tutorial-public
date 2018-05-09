func Open(name string) (file *File, err error) {
    return OpenFile(name, O_RDONLY, 0)
}

func (f *File) Chdir() error {
    if f == nil {
        return ErrInvalid
    }
    if e := syscall.Fchdir(f.fd); e != nil {
        return &PathError{"chdir", f.name, e}
    }
    return nil
}

func epipecheck(file *File, e error) {
    if e == syscall.EPIPE {
        if atomic.AddInt32(&file.nepipe, 1) >= 10 {
            sigpipe()
        }
    } else {
        atomic.StoreInt32(&file.nepipe, 0)
    }
}
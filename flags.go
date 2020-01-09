package main




const (
	DIRTY  = 1
	DELETED    = 2
	BLOCKED    = 4
	EXTERNAL = 8
	SUPERUSER = 16
	NEW = 32
)


type Flag struct {
	val uint
}


func NewFlag() *Flag {
	f := Flag{}
	f.val = 0
	return &f
}

func (f *Flag) Is(flag uint)  bool {
	return f.val & flag != 0
}


func (f *Flag) Set(flag uint)   {
	f.val  |=  flag
}

func (f *Flag) Clear(flag uint)   {
	f.val   &^=   flag
}

func (f *Flag) Get() uint  {
	return f.val
}

func (f *Flag) Init(flag uint)   {
	f.val = flag
}





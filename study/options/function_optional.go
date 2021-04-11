/**
 * @link https://halls-of-valhalla.org/beta/articles/functional-options-pattern-in-go,54/
 * 作为Golang开发人员，您会遇到的许多问题之一就是尝试使函数的参数成为可选参数。这是一个非常常见的用例，
 * 您有一些对象可以使用一些基本的默认设置即开即用地工作，并且您有时可能希望提供一些更详细的配置。
 * 在许多语言中，这很容易。在C系列语言中，您可以提供具有不同数量参数的同一函数的多个版本；在PHP之类的语言中，可以为参数提供默认值，并在调用方法时将其忽略。
 * 但是在Golang中，您不能做任何一个。那么，如何创建一个具有一些其他配置的功能，用户可以指定是否需要这些功能，但只有他们愿意时可以指定？
 * 解决方案：Functional Options Pattern 选项模式
 */

package options

type Connection struct {
}

type StuffClient interface {
	DoStuff() error
}

type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}

func (s *stuffClient) DoStuff() error {
	return nil
}

// ===================================================================================================================
// pass timeout and retries
func NewStuffClientWithValue(conn Connection, timeout int, retries int) StuffClient {
	return &stuffClient{
		conn:    conn,
		timeout: timeout,
		retries: retries,
	}
}

// ===================================================================================================================
const (
	DefaultTimeout = 5
	DefaultRetries = 3
)

// use default timeout and retries
func NewStuffClient(conn Connection) StuffClient {
	return &stuffClient{
		conn:    conn,
		timeout: DefaultTimeout,
		retries: DefaultRetries,
	}
}

// ===================================================================================================================
type StuffClientOptions struct {
	Timeout int
	Retries int
}

// use options struct
func NewStuffClientWithOptionsStruct(conn Connection, options StuffClientOptions) StuffClient {
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}

// ===================================================================================================================
// closure 闭包函数
type StuffClientOption func(options *StuffClientOptions)

func WithTimeout(t int) StuffClientOption {
	return func(options *StuffClientOptions) {
		options.Timeout = t
	}
}

func WithRetries(r int) StuffClientOption {
	return func(options *StuffClientOptions) {
		options.Retries = r
	}
}

var defaultStuffClientOptions = StuffClientOptions{
	Timeout: DefaultTimeout,
	Retries: DefaultRetries,
}

// use options
// default: NewStuffClientWithOptions(Connection{})
// optional: NewStuffClientWithOptions(Connection{}, WithTimeout(77), WithRetries(66))
func NewStuffClientWithOptions(conn Connection, opts ...StuffClientOption) StuffClient {
	options := defaultStuffClientOptions
	for _, o := range opts {
		o(&options)
	}
	return &stuffClient{
		conn:    conn,
		timeout: options.Timeout,
		retries: options.Retries,
	}
}

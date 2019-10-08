package emsv1

/*
#cgo darwin CFLAGS: -I.
#cgo darwin CFLAGS: -I/home/include/tibems
#cgo darwin LDFLAGS: -L/home/lib -ltibems64

#include <tibems.h>
tibemsDestination castToDestination(tibemsTemporaryQueue queue) {
  return (tibemsDestination)queue;
}
tibems_bool castToBool(int value) {
	return (tibems_bool)value;
}
tibems_long castToLong(int value) {
  return (tibems_long)value;
}
tibems_int castToInt(int value) {
  return (tibems_int)value;
}

*/
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	/* test */
	TIBEMS_SESSION_TRANSACTED = 0
	/* test */
	TIBEMS_AUTO_ACKNOWLEDGE = 1
	/* test */
	TIBEMS_CLIENT_ACKNOWLEDGE = 2
	/* test */
	TIBEMS_DUPS_OK_ACKNOWLEDGE = 3
	/* test */
	TIBEMS_NO_ACKNOWLEDGE = 22 /* Extensions */
	/* test */
	TIBEMS_EXPLICIT_CLIENT_ACKNOWLEDGE = 23
	/* test */
	TIBEMS_EXPLICIT_CLIENT_DUPS_OK_ACKNOWLEDGE = 24
) //tibemsAcknowledgeMode

const (
	/* test */
	TIBEMS_FALSE = 0
	/* test */
	TIBEMS_TRUE = 1
) //tibems_bool

const (
	/* test */
	TIBEMS_UNKNOWN = 0
	/* test */
	TIBEMS_QUEUE = 1
	/* test */
	TIBEMS_TOPIC = 2
	/* test */
	TIBEMS_DEST_UNDEFINED = 256
) //tibemsDestinationType

const (
	/* test */
	TIBEMS_NON_PERSISTENT = 1
	/* test */
	TIBEMS_PERSISTENT = 2
	/* test */
	TIBEMS_RELIABLE = 22 /* Extension */
) //tibemsDeliveryMode

const (
	/* test */
	NPSEND_CHECK_DEFAULT = 0
	/* test */
	NPSEND_CHECK_ALWAYS = 1
	/* test */
	NPSEND_CHECK_NEVER = 2
	/* test */
	NPSEND_CHECK_TEMP_DEST = 3
	/* test */
	NPSEND_CHECK_AUTH = 4
	/* test */
	NPSEND_CHECK_TEMP_AUTH = 5
) //tibemsNpCheckMode

const (
	TIBEMS_MESSAGE_UNKNOWN   = 0
	TIBEMS_MESSAGE           = 1
	TIBEMS_BYTES_MESSAGE     = 2
	TIBEMS_MAP_MESSAGE       = 3
	TIBEMS_OBJECT_MESSAGE    = 4
	TIBEMS_STREAM_MESSAGE    = 5
	TIBEMS_TEXT_MESSAGE      = 6
	TIBEMS_MESSAGE_UNDEFINED = 256
) //tibemsMsgType

const (
	TIBEMS_OK = 0

	TIBEMS_ILLEGAL_STATE       = 1
	TIBEMS_INVALID_CLIENT_ID   = 2
	TIBEMS_INVALID_DESTINATION = 3
	TIBEMS_INVALID_SELECTOR    = 4

	TIBEMS_EXCEPTION          = 5
	TIBEMS_SECURITY_EXCEPTION = 6

	TIBEMS_MSG_EOF = 7

	TIBEMS_MSG_NOT_READABLE  = 9
	TIBEMS_MSG_NOT_WRITEABLE = 10

	TIBEMS_SERVER_NOT_CONNECTED = 11
	TIBEMS_VERSION_MISMATCH     = 12
	TIBEMS_SUBJECT_COLLISION    = 13

	TIBEMS_INVALID_PROTOCOL = 15
	TIBEMS_INVALID_HOSTNAME = 17
	TIBEMS_INVALID_PORT     = 18
	TIBEMS_NO_MEMORY        = 19
	TIBEMS_INVALID_ARG      = 20

	TIBEMS_SERVER_LIMIT = 21

	TIBEMS_MSG_DUPLICATE = 22

	TIBEMS_SERVER_DISCONNECTED = 23
	TIBEMS_SERVER_RECONNECTING = 24

	TIBEMS_NOT_PERMITTED = 27

	TIBEMS_SERVER_RECONNECTED = 28

	TIBEMS_INVALID_NAME      = 30
	TIBEMS_INVALID_TYPE      = 31
	TIBEMS_INVALID_SIZE      = 32
	TIBEMS_INVALID_COUNT     = 33
	TIBEMS_NOT_FOUND         = 35
	TIBEMS_ID_IN_USE         = 36
	TIBEMS_ID_CONFLICT       = 37
	TIBEMS_CONVERSION_FAILED = 38

	TIBEMS_INVALID_MSG      = 42
	TIBEMS_INVALID_FIELD    = 43
	TIBEMS_INVALID_INSTANCE = 44
	TIBEMS_CORRUPT_MSG      = 45

	TIBEMS_PRODUCER_FAILED = 47

	TIBEMS_TIMEOUT                    = 50
	TIBEMS_INTR                       = 51
	TIBEMS_DESTINATION_LIMIT_EXCEEDED = 52
	TIBEMS_MEM_LIMIT_EXCEEDED         = 53
	TIBEMS_USER_INTR                  = 54

	TIBEMS_INVALID_QUEUE_GROUP   = 63
	TIBEMS_INVALID_TIME_INTERVAL = 64
	TIBEMS_INVALID_IO_SOURCE     = 65
	TIBEMS_INVALID_IO_CONDITION  = 66
	TIBEMS_SOCKET_LIMIT          = 67

	TIBEMS_OS_ERROR = 68

	TIBEMS_WOULD_BLOCK = 69

	TIBEMS_INSUFFICIENT_BUFFER = 70

	TIBEMS_EOF            = 71
	TIBEMS_INVALID_FILE   = 72
	TIBEMS_FILE_NOT_FOUND = 73
	TIBEMS_IO_FAILED      = 74

	TIBEMS_NOT_FILE_OWNER = 80

	TIBEMS_ALREADY_EXISTS = 91

	TIBEMS_INVALID_CONNECTION = 100
	TIBEMS_INVALID_SESSION    = 101
	TIBEMS_INVALID_CONSUMER   = 102
	TIBEMS_INVALID_PRODUCER   = 103
	TIBEMS_INVALID_USER       = 104
	TIBEMS_INVALID_GROUP      = 105

	TIBEMS_TRANSACTION_FAILED   = 106
	TIBEMS_TRANSACTION_ROLLBACK = 107
	TIBEMS_TRANSACTION_RETRY    = 108

	TIBEMS_INVALID_XARESOURCE = 109

	TIBEMS_FT_SERVER_LACKS_TRANSACTION = 110

	TIBEMS_LDAP_ERROR         = 120
	TIBEMS_INVALID_PROXY_USER = 121

	/* SSL related errors */
	TIBEMS_INVALID_CERT         = 150
	TIBEMS_INVALID_CERT_NOT_YET = 151
	TIBEMS_INVALID_CERT_EXPIRED = 152
	TIBEMS_INVALID_CERT_DATA    = 153
	TIBEMS_ALGORITHM_ERROR      = 154
	TIBEMS_SSL_ERROR            = 155
	TIBEMS_INVALID_PRIVATE_KEY  = 156
	TIBEMS_INVALID_ENCODING     = 157
	TIBEMS_NOT_ENOUGH_RANDOM    = 158
	TIBEMS_INVALID_CRL_DATA     = 159
	TIBEMS_CRL_OFF              = 160
	TIBEMS_EMPTY_CRL            = 161

	TIBEMS_NOT_INITIALIZED    = 200
	TIBEMS_INIT_FAILURE       = 201
	TIBEMS_ARG_CONFLICT       = 202
	TIBEMS_SERVICE_NOT_FOUND  = 210
	TIBEMS_INVALID_CALLBACK   = 211
	TIBEMS_INVALID_QUEUE      = 212
	TIBEMS_INVALID_EVENT      = 213
	TIBEMS_INVALID_SUBJECT    = 214
	TIBEMS_INVALID_DISPATCHER = 215

	/* JVM related errors */
	TIBEMS_JNI_EXCEPTION = 230
	TIBEMS_JNI_ERR       = 231
	TIBEMS_JNI_EDETACHED = 232
	TIBEMS_JNI_EVERSION  = 233
	TIBEMS_JNI_EEXIST    = 235
	TIBEMS_JNI_EINVAL    = 236

	TIBEMS_NO_MEMORY_FOR_OBJECT = 237

	TIBEMS_UFO_CONNECTION_FAILURE = 240

	TIBEMS_NOT_IMPLEMENTED = 255
)

const (
	disconnected uint32 = iota
	connected
)

type ClientOptions struct {
	serverUrl url.URL
	username  string
	password  string
}

func NewClientOptions() *ClientOptions {
	o := &ClientOptions{
		username: "",
		password: "",
	}

	return o
}

func (o *ClientOptions) SetServerUrl(p string) *ClientOptions {

	url, err := url.Parse(p)
	if err == nil {
		o.serverUrl = *url
	}
	return o
}

func (o *ClientOptions) SetUsername(p string) *ClientOptions {
	o.username = p
	return o
}

func (o *ClientOptions) SetPassword(p string) *ClientOptions {
	o.password = p
	return o
}

func (o *ClientOptions) GetServerUrl() url.URL {
	return o.serverUrl
}

func (o *ClientOptions) GetUsername() string {
	return o.username
}

func (o *ClientOptions) GetPassword() string {
	return o.password
}

/* test */
type IClient interface {
	IsConnected() bool
	Connect() error
	Disconnect() error
	Send(destination string, message string, deliveryDelay int, deliveryMode string, expiration int) error
	SendReceive(destination string, message string, deliveryMode string, expiration int) (string, error)
	Receive(destination string) (string, error)
}

/* test */
type Client struct {
	conn         C.tibemsConnection
	cf           C.tibemsConnectionFactory
	errorContext C.tibemsErrorContext
	status       uint32
	options      ClientOptions
	sync.RWMutex
}

func NewClient(o *ClientOptions) IClient {

	c := &Client{}
	c.options = *o
	c.status = disconnected

	return c
}

func (c *Client) IsConnected() bool {

	c.RLock()
	defer c.RUnlock()

	return c.status == connected

}
func (c *Client) Connect() error {

	c.RLock()
	defer c.RUnlock()

	status := C.tibemsErrorContext_Create(&c.errorContext)

	if status != TIBEMS_OK {
		return errors.New("failed to create error context")
	}

	c.cf = C.tibemsConnectionFactory_Create()

	url := c.options.GetServerUrl()

	status = C.tibemsConnectionFactory_SetServerURL(c.cf, C.CString(url.String()))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// create the connection
	status = C.tibemsConnectionFactory_CreateConnection(c.cf, &c.conn, C.CString(c.options.username), C.CString(c.options.password))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// start the connection
	status = C.tibemsConnection_Start(c.conn)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	c.setConnected(connected)

	return nil
}

func (c *Client) Disconnect() error {

	c.RLock()
	defer c.RUnlock()

	if c.IsConnected() {

		status := C.tibemsConnection_Stop(c.conn)
		if status != TIBEMS_OK {
			return errors.New("failed to stop connection")
		}

		// close the connection
		status = C.tibemsConnection_Close(c.conn)
		if status != TIBEMS_OK {
			return errors.New("failed to close connection")
		}

		c.setConnected(disconnected)
	}

	return nil
}

func (c *Client) SendReceive(destination string, message string, deliveryMode string, expiration int) (string, error) {
	var dest C.tibemsDestination
	var session C.tibemsSession
	var requestor C.tibemsMsgRequestor
	var reqMsg C.tibemsMsg
	var repMsg C.tibemsMsg

	var msg C.tibemsTextMsg

	// create the destination
	status := C.tibemsDestination_Create(&dest, TIBEMS_QUEUE, C.CString(destination))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the session
	status = C.tibemsConnection_CreateSession(c.conn, &session, TIBEMS_FALSE, TIBEMS_AUTO_ACKNOWLEDGE)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the requestor
	status = C.tibemsMsgRequestor_Create(session, &requestor, dest)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the request message
	status = C.tibemsMsg_Create(&reqMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the message
	status = C.tibemsTextMsg_Create(&msg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// set message delivery mode
	var emsDeliveryMode = TIBEMS_NON_PERSISTENT
	if strings.ToLower(deliveryMode) == "persistent" {
		emsDeliveryMode = TIBEMS_PERSISTENT
	} else if strings.ToLower(deliveryMode) == "non_persistent" {
		emsDeliveryMode = TIBEMS_NON_PERSISTENT
	} else if strings.ToLower(deliveryMode) == "reliable" {
		emsDeliveryMode = TIBEMS_RELIABLE
	}

	status = C.tibemsMsg_SetDeliveryMode(msg, C.tibemsDeliveryMode(emsDeliveryMode))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// set message expiration
	status = C.tibemsMsg_SetExpiration(msg, C.castToLong(C.int(expiration)))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the reply message
	status = C.tibemsMsg_Create(&repMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// set the message text
	status = C.tibemsTextMsg_SetText(msg, C.CString(message))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// send a request message; wait for a reply
	status = C.tibemsMsgRequestor_Request(requestor, msg, &repMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// Get the string data from the reply text message

	var buf *C.char
	buf = (*C.char)(C.calloc(32768, 1))
	defer C.free(unsafe.Pointer(buf))

	status = C.tibemsTextMsg_GetText(repMsg, &buf)

	replyMessageText := C.GoString(buf)

	fmt.Println("Received JMS Reply Text Message = " + replyMessageText)

	// destroy the request message
	status = C.tibemsMsg_Destroy(reqMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// destroy the requestor
	status = C.tibemsMsgRequestor_Close(requestor)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// destroy the session
	status = C.tibemsSession_Close(session)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// destroy the destination
	status = C.tibemsDestination_Destroy(dest)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	return replyMessageText, nil

}

func (c *Client) Receive(destination string) (string, error) {

	var dest C.tibemsDestination
	var session C.tibemsSession
	var msgConsumer C.tibemsMsgConsumer
	var msg C.tibemsMsg
	var msgType C.tibemsMsgType
	var msgTypeName string

	// create the destination
	status := C.tibemsDestination_Create(&dest, TIBEMS_QUEUE, C.CString(destination))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the session
	status = C.tibemsConnection_CreateSession(c.conn, &session, TIBEMS_FALSE, TIBEMS_AUTO_ACKNOWLEDGE)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// create the consumer
	status = C.tibemsSession_CreateConsumer(session, &msgConsumer, dest, nil, TIBEMS_FALSE)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// start the connection
	status = C.tibemsConnection_Start(c.conn)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	status = C.tibemsMsgConsumer_Receive(msgConsumer, &msg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// Check message type
	status = C.tibemsMsg_GetBodyType(msg, &msgType)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	switch msgType {
	case TIBEMS_MESSAGE:
		msgTypeName = "MESSAGE"
		break
	case TIBEMS_TEXT_MESSAGE:
		msgTypeName = "TEXT"
		break
	case TIBEMS_BYTES_MESSAGE:
		msgTypeName = "BYTES"
		break
	case TIBEMS_OBJECT_MESSAGE:
		msgTypeName = "OBJECT"
		break
	case TIBEMS_STREAM_MESSAGE:
		msgTypeName = "STREAM"
		break
	case TIBEMS_MAP_MESSAGE:
		msgTypeName = "MAP"
		break
	default:
		msgTypeName = "UNKNOWN"
		break
	}

	if msgType != TIBEMS_TEXT_MESSAGE {
		return "", errors.New("Unable to process message type " + msgTypeName)
	}

	// Get the string data from the reply text message

	var buf *C.char
	buf = (*C.char)(C.calloc(32768, 1))
	defer C.free(unsafe.Pointer(buf))

	status = C.tibemsTextMsg_GetText(msg, &buf)

	messageText := C.GoString(buf)

	//fmt.Println("Received JMS Text Message = "+ messageText)

	// destroy the message
	status = C.tibemsMsg_Destroy(msg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// destroy the session
	status = C.tibemsSession_Close(session)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	// destroy the destination
	status = C.tibemsDestination_Destroy(dest)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return "", errors.New(e)
	}

	return messageText, nil
}

func (c *Client) Send(destination string, message string, deliveryDelay int, deliveryMode string, expiration int) error {

	var dest C.tibemsDestination
	var session C.tibemsSession
	var msgProducer C.tibemsMsgProducer
	var txtMsg C.tibemsTextMsg

	// create the destination
	status := C.tibemsDestination_Create(&dest, TIBEMS_QUEUE, C.CString(destination))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// create the session
	status = C.tibemsConnection_CreateSession(c.conn, &session, TIBEMS_FALSE, TIBEMS_AUTO_ACKNOWLEDGE)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// create the producer
	status = C.tibemsSession_CreateProducer(session, &msgProducer, dest)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	status = C.tibemsMsgProducer_SetDeliveryDelay(msgProducer, C.castToLong(C.int(deliveryDelay)))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	var emsDeliveryMode = TIBEMS_NON_PERSISTENT
	if strings.ToLower(deliveryMode) == "persistent" {
		emsDeliveryMode = TIBEMS_PERSISTENT
	} else if strings.ToLower(deliveryMode) == "non_persistent" {
		emsDeliveryMode = TIBEMS_NON_PERSISTENT
	} else if strings.ToLower(deliveryMode) == "reliable" {
		emsDeliveryMode = TIBEMS_RELIABLE
	}

	status = C.tibemsMsgProducer_SetDeliveryMode(msgProducer, C.castToInt(C.int(emsDeliveryMode)))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	status = C.tibemsMsgProducer_SetTimeToLive(msgProducer, C.castToLong(C.int(expiration)))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// create the message
	status = C.tibemsTextMsg_Create(&txtMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// set the message text
	status = C.tibemsTextMsg_SetText(txtMsg, C.CString(message))
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// publish the message
	status = C.tibemsMsgProducer_Send(msgProducer, txtMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// destroy the message
	status = C.tibemsMsg_Destroy(txtMsg)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// destroy the producer
	status = C.tibemsMsgProducer_Close(msgProducer)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// destroy the session
	status = C.tibemsSession_Close(session)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	// destroy the destination
	status = C.tibemsDestination_Destroy(dest)
	if status != TIBEMS_OK {
		e, _ := c.getErrorContext()
		return errors.New(e)
	}

	return nil
}

func (c *Client) connectionStatus() uint32 {
	c.RLock()
	defer c.RUnlock()
	status := atomic.LoadUint32(&c.status)
	return status
}

func (c *Client) setConnected(status uint32) {
	c.RLock()
	defer c.RUnlock()
	atomic.StoreUint32(&c.status, status)
}

func (c *Client) getErrorContext() (string, string) {

	var errorString, stackTrace = "", ""
	var buf *C.char
	defer C.free(unsafe.Pointer(buf))

	C.tibemsErrorContext_GetLastErrorString(c.errorContext, &buf)
	errorString = C.GoString(buf)

	C.tibemsErrorContext_GetLastErrorStackTrace(c.errorContext, &buf)
	stackTrace = C.GoString(buf)

	return errorString, stackTrace

}

const (
	ivContent       = "content"
	ivDestination   = "destination"
	ivServerURL     = "serverUrl"
	ivUser          = "user"
	ivPassword      = "password"
	ivDeliveryDelay = "deliveryDelay"
	ivDeliveryMode  = "deliveryMode"
	ivExpiration    = "expiration"
	ivTracing       = "tracing"
	ivExchangeMode  = "exchangeMode"

	ovResponse = "response"
	ovTracing  = "tracing"
)

var (
	errorDestinationIsNotAString         = errors.New("destination is not a string")
	errorInvalidEmptyDestinationToSendTo = errors.New("invalid empty destination to send to")
	errorDeliveryDelayIsNotANumber       = errors.New("delivery delay is not a number")
	errorDeliveryModeIsNotAString        = errors.New("delivery mode is not a string")
	errorExpirationIsNotANumber          = errors.New("expiration is not a number")
	errorExchangeModeIsNotAString        = errors.New("exchange mode is not a string")
)

var log = logger.GetLogger("activity-tibco-ems")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	log.Debugf("********************************************** line 1")

	//var span opentracing.Span
	if tracing := context.GetInput(ivTracing); tracing != nil {
		//span = opentracing.SpanFromContext(tracing.(ctx.Context))
		log.Debugf("*********************************************** error 1")
	}

	/*if span != nil {
		span = opentracing.StartSpan(
			context.TaskName(),
			opentracing.ChildOf(span.Context()))
		context.SetOutput(ovTracing, opentracing.ContextWithSpan(ctx.Background(), span))
		defer span.Finish()
	}*/

	setTag := func(key string, value interface{}) {
		/*if span != nil {
			span.SetTag(key, value)
		}*/
		log.Debugf("********************************** key %v ", key)
	}

	logError := func(format string, a ...interface{}) {
		str := fmt.Sprintf(format, a...)
		setTag("error", str)
		log.Error(str)
	}

	opts := NewClientOptions()

	if serverURL, ok := context.GetInput(ivServerURL).(string); ok {
		setTag("serverUrl", serverURL)
		opts.SetServerUrl(serverURL)
		log.Debugf("************************************************ server %v ", serverURL)
	}

	if user, ok := context.GetInput(ivUser).(string); ok {
		opts.SetUsername(user)
		log.Debugf("************************************************  user %v ", user)
	}

	if password, ok := context.GetInput(ivPassword).(string); ok {
		opts.SetPassword(password)
	} else {
		opts.SetPassword("")
	}

	client := NewClient(opts)

	err = client.Connect()
	if err != nil {
		log.Debugf("************************************************  connect error  %v ", err)
		logError("Connection to EMS Server failed %v", err.Error())
	}
	defer client.Disconnect()

	content := ""
	switch v := context.GetInput(ivContent).(type) {
	case int, int64, float64, bool, json.Number:
		content = fmt.Sprintf("%v", v)
	case string:
		content = v
	default:
		var data []byte
		data, err = json.Marshal(v)
		if err != nil {
			logError("Invalid content %v", err)
			break
		}
		content = string(data)
	}
	setTag("content", content)
	log.Debugf("************************************************  content %v ", content)
	destination, ok := context.GetInput(ivDestination).(string)
	if !ok {
		logError(errorDestinationIsNotAString.Error())
		return false, errorDestinationIsNotAString
	}
	if len(destination) == 0 {
		logError(errorInvalidEmptyDestinationToSendTo.Error())
		return false, errorInvalidEmptyDestinationToSendTo
	}
	setTag("destination", destination)

	deliveryDelay, ok := context.GetInput(ivDeliveryDelay).(int)
	if !ok {
		logError(errorDeliveryDelayIsNotANumber.Error())
		return false, errorDeliveryDelayIsNotANumber
	}

	expiration, ok := context.GetInput(ivExpiration).(int)
	if !ok {
		logError(errorExpirationIsNotANumber.Error())
		return false, errorExpirationIsNotANumber
	}

	deliveryMode, ok := context.GetInput(ivDeliveryMode).(string)
	if !ok {
		logError(errorDeliveryModeIsNotAString.Error())
		return false, errorDeliveryModeIsNotAString
	}

	exchangeMode, ok := context.GetInput(ivExchangeMode).(string)
	if !ok {
		logError(errorExchangeModeIsNotAString.Error())
		return false, errorDestinationIsNotAString
	}
	log.Debugf("************************************************ other ")

	if exchangeMode == "send-only" {
		fmt.Printf("************************************************  sending ")
		err = client.Send(destination, content, deliveryDelay, deliveryMode, expiration)
		if err != nil {
			log.Debugf("************************************************  sending err %v ", err)
			logError("Timeout occurred while trying to send to destination '%s'", destination)
			return false, err
		}
	} else {
		response, err := client.SendReceive(destination, content, deliveryMode, expiration)
		if err != nil {
			logError("Timeout occurred while trying to send to destination '%s'", destination)
			return false, err
		}

		log.Debugf("Response payload: %s", response)
		context.SetOutput(ovResponse, response)

	}

	return true, nil
}

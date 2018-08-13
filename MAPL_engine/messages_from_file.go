package MAPL_engine


import (
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
	"time"
)


func YamlReadMessageAttributes(yamlString string) MessageAttributes {

	var messageAttributes MessageAttributes
	err := yaml.Unmarshal([]byte(yamlString), &messageAttributes)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("---values found:\n%+v\n\n", rule)

	flag, outputString := IsNumberOfFieldsEqual(messageAttributes, yamlString)

	if flag == false {
		panic("number of fields in rules does not match number of fields in yaml file:\n" + outputString)
	}

	AddResourceType(&messageAttributes)

	return messageAttributes
}

func YamlReadMessagesFromString(yamlString string) Messages {

	var messages Messages
	err := yaml.Unmarshal([]byte(yamlString), &messages)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("---values found:\n%+v\n\n", rule)

	addResourceTypeToMessages(&messages)
	addTimeInfoToMessages(&messages)


	flag, outputString := IsNumberOfFieldsEqual(messages, yamlString)
	if flag == false {
		panic("number of fields in rules does not match number of fields in yaml file:\n" + outputString)
	}

	return messages
}

func YamlReadMessagesFromFile(filename string) Messages {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	messages := YamlReadMessagesFromString(string(data))
	return messages
}

func AddResourceType(message *MessageAttributes){
	// add resource_type by the resource_protocol
	// we have resource_type to allow for several types per one protocol.
	//
	message.ContextType=""
	switch message.ContextProtocol{ // these are the only protocols we currently support
	case "HTTP","http":
		message.ContextType = "httpPath"
	case "TCP","tcp":
		message.ContextType = "port"
	}
}

func addResourceTypeToMessages(messages *Messages) {
	// add resource_type by the resource_protocol
	// we have resource_type to allow for several types per one protocol.

	for i, _ := range(messages.Messages) {
		AddResourceType(&messages.Messages[i])
	}

}

func addTimeInfoToMessage(message *MessageAttributes) {
	//
	// extract timestamp info
	//

	//t, err := time.Parse(time.RFC3339,"2018-07-29T14:30:00-07:00")
	t, err := time.Parse(time.RFC3339,message.RequestTime)
	if err!=nil{
		panic("error in parsing message time")
	}

	nanosecondsFromMidnight := float64(((t.Hour()*60+t.Minute())*60+t.Second())*1e9+t.Nanosecond())

	message.RequestTimeSecondsFromMidnightUTC = nanosecondsFromMidnight/1e9
	message.RequestTimeMinutesFromMidnightUTC = nanosecondsFromMidnight/1e9/60
	message.RequestTimeHoursFromMidnightUTC = nanosecondsFromMidnight/1e9/60/60
}

func addTimeInfoToMessages(messages *Messages) {
	// add resource_type by the resource_protocol
	// we have resource_type to allow for several types per one protocol.

	for i, _ := range(messages.Messages) {
		addTimeInfoToMessage(&messages.Messages[i])
	}
}
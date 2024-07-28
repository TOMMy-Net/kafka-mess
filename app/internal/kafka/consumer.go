package kafka


func ReadMessage() (string, error) {

		mess, err := Broker.ReadMessage(1024)
		if err != nil {
			return "", err
		}
		return string(mess.Value), nil
	
}

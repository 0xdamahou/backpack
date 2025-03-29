package authenticated

func (c *BackpackClient) GetOpenPositions() (Positions, error) {
	endpoint := "api/v1/position"
	instruction := "positionQuery"
	var result Positions
	err := c.DoGet(endpoint, instruction, "", &result)
	return result, err

}

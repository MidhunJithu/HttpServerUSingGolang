package utils

import "example/httpserver/models"

func CopyToStructuredEventStruct(E *models.Event) models.StrcturedEvent {
	return models.StrcturedEvent{
		Event:      E.Event,
		Trigger:    E.Trigger,
		Id:         E.Id,
		MessageId:  E.MessageId,
		Title:      E.Title,
		Page:       E.Page,
		Language:   E.Language,
		ScreenSize: E.ScreenSize,
		Attributes: map[string]interface{}{
			E.AttrKey1: map[string]string{"value": E.AttrValue1, "type": E.AttrType1},
			E.AttrKey2: map[string]string{"value": E.AttrValue2, "type": E.AttrType2},
		},
		Traits: map[string]interface{}{
			E.UserTraitKey1: map[string]string{"value": E.UserTraitVal1, "type": E.UserTraitTYpe1},
			E.UserTraitKey2: map[string]string{"value": E.UserTraitVal2, "type": E.UserTraitTYpe2},
			E.UserTraitKey3: map[string]string{"value": E.UserTraitVal3, "type": E.UserTraitTYpe3},
		},
	}
}

package turnsignal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignal(t *testing.T) {
	count, setCount := CreateSignal(0)
	allExpected := []int{0, 5, 10}
	CreateEffect(func() {
		assert.NotEmpty(t, allExpected)
		expected := allExpected[0]
		allExpected = allExpected[1:]
		assert.Equal(t, expected, count())
	})

	setCount(5)
	setCount(10)

	assert.Empty(t, allExpected)
	assert.Empty(t, effectContextStack)

}

func TestConditionalRendering(t *testing.T) {

	firstName, _ := CreateSignal("John")
	lastName, setLastName := CreateSignal("Smith")
	showFullName, setShowFullName := CreateSignal(true)

	displayName := CreateMemo(func() string {
		if !showFullName() {
			return firstName()
		}
		return fmt.Sprintf("%s %s", firstName(), lastName())
	}, "")
	allConditionalExpected := []string{"John Smith", "John", "John Legend"}
	CreateEffect(func() {
		assert.NotEmpty(t, allConditionalExpected)
		conditionalExpected := allConditionalExpected[0]
		allConditionalExpected = allConditionalExpected[1:]
		actualDisplayName := displayName()
		assert.Equal(t, conditionalExpected, actualDisplayName)
	})

	fullName := func() string {
		return fmt.Sprintf("%s %s", firstName(), lastName())
	}
	allFullExpected := []string{"John Smith", "John Legend", "John Legend"}
	CreateEffect(func() {
		assert.NotEmpty(t, allFullExpected)
		fullExpected := allFullExpected[0]
		allFullExpected = allFullExpected[1:]
		actualFullName := fullName()
		assert.Equal(t, fullExpected, actualFullName)
	})

	setShowFullName(false)
	setLastName("Legend")
	setShowFullName(true)

	assert.Empty(t, effectContextStack)
	assert.Empty(t, allConditionalExpected)
}

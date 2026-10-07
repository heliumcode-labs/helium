package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassifyUsage(t *testing.T) {
	assert.Equal(t, usageHealthy, classifyUsage(0))
	assert.Equal(t, usageHealthy, classifyUsage(69.9))
	assert.Equal(t, usageBusy, classifyUsage(70))
	assert.Equal(t, usageBusy, classifyUsage(89.9))
	assert.Equal(t, usageCritical, classifyUsage(90))
	assert.Equal(t, usageCritical, classifyUsage(100))
}

func TestFormatTokensAndCost(t *testing.T) {
	// The percentage is always part of the string so it can be tracked before
	// the warning threshold is reached.
	assert.Equal(t, "Context: 1.2K (12%), Cost: $0.10",
		formatTokensAndCost(1200, 10_000, 0.1))

	assert.Equal(t, "Context: 420 (0%), Cost: $0.00",
		formatTokensAndCost(420, 1_000_000, 0))

	// No context window configured: fall back to 0% instead of dividing by zero.
	assert.Equal(t, "Context: 1K (0%), Cost: $1.00",
		formatTokensAndCost(1000, 0, 1))

	assert.Equal(t, "Context: 2.5M (50%), Cost: $12.34",
		formatTokensAndCost(2_500_000, 5_000_000, 12.34))
}

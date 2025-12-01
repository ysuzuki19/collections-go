package frame_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ysuzuki19/collections-go/traceback/internal/frame"
)

func TestCapture(t *testing.T) {
	require := require.New(t)

	{
		// testdoc begin Capture
		frames := frame.Capture(0)
		// testdoc end
		require.Equal(4, frames.Len()) // runtime>testing>TestCapture>frame.Capture
	}

	{
		frames := frame.Capture(1)
		require.Equal(3, frames.Len()) // runtime>testing>TestCapture
	}

	{
		frames := func() frame.Frames {
			return frame.Capture(0)
		}()
		require.Equal(5, frames.Len()) // runtime>testing>func1>TestCapture>frame.Capture
	}

	{
		frames := func() frame.Frames {
			return frame.Capture(1)
		}()
		require.Equal(4, frames.Len()) // runtime>testing>func1>TestCapture
	}
}

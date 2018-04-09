package rules

import (
	"testing"

	"github.com/battlesnakeio/engine/controller/pb"
	"github.com/stretchr/testify/require"
)

func TestDeathCauseStarvation(t *testing.T) {
	updates := checkForDeath(20, 20, &pb.GameTick{
		Turn: 3,
		Snakes: []*pb.Snake{
			&pb.Snake{
				Health: 0,
			},
		},
	})
	require.Len(t, updates, 1)
	require.Equal(t, DeathCauseStarvation, updates[0].Death.Cause)
	require.Equal(t, int64(3), updates[0].Death.Turn)
}

func TestDeathCauseWallCollision(t *testing.T) {
	points := []*pb.Point{

		{X: -1, Y: 1},
		{X: 20, Y: 1},
		{X: 1, Y: -1},
		{X: 1, Y: 20},
	}
	for _, p := range points {
		updates := checkForDeath(20, 20, &pb.GameTick{
			Turn: 3,
			Snakes: []*pb.Snake{
				&pb.Snake{
					Health: 45,
					Body:   []*pb.Point{p},
				},
			},
		})
		require.Len(t, updates, 1)
		require.Equal(t, DeathCauseWallCollision, updates[0].Death.Cause)
		require.Equal(t, int64(3), updates[0].Death.Turn)
	}
}

func TestDeathCauseSnakeCollision(t *testing.T) {
	updates := checkForDeath(20, 20, &pb.GameTick{
		Turn: 3,
		Snakes: []*pb.Snake{
			&pb.Snake{
				ID:     "1",
				Health: 45,
				Body: []*pb.Point{
					{X: 5, Y: 5},
				},
			},
			&pb.Snake{
				ID:     "2",
				Health: 56,
				Body: []*pb.Point{
					{X: 6, Y: 5},
					{X: 5, Y: 5},
				},
			},
		},
	})
	require.Len(t, updates, 1)
	require.Equal(t, DeathCauseSnakeCollision, updates[0].Death.Cause)
	require.Equal(t, int64(3), updates[0].Death.Turn)
}

func TestDeathCauseHeadToHeadCollision(t *testing.T) {
	updates := checkForDeath(20, 20, &pb.GameTick{
		Turn: 3,
		Snakes: []*pb.Snake{
			&pb.Snake{
				ID:     "1",
				Health: 45,
				Body: []*pb.Point{
					{X: 6, Y: 5},
					{X: 5, Y: 5},
				},
			},
			&pb.Snake{
				ID:     "2",
				Health: 56,
				Body: []*pb.Point{
					{X: 6, Y: 5},
					{X: 5, Y: 5},
				},
			},
		},
	})
	require.Len(t, updates, 2)
	require.Equal(t, DeathCauseHeadToHeadCollision, updates[0].Death.Cause)
	require.Equal(t, int64(3), updates[0].Death.Turn)
	require.Equal(t, DeathCauseHeadToHeadCollision, updates[1].Death.Cause)
	require.Equal(t, int64(3), updates[1].Death.Turn)
}
// +build integration

package sqs_test

import (
	"testing"

	"github.com/datacratic/aws-sdk-go/aws"
	"github.com/datacratic/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func TestFlattenedTraits(t *testing.T) {
	s := sqs.New(nil)
	_, err := s.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
		QueueURL: aws.String("QUEUE"),
		Entries: []*sqs.DeleteMessageBatchRequestEntry{
			&sqs.DeleteMessageBatchRequestEntry{
				ID:            aws.String("TEST"),
				ReceiptHandle: aws.String("RECEIPT"),
			},
		},
	})

	assert.Error(t, err)
	assert.EqualError(t, err, "The address QUEUE is not valid for this endpoint.")
}

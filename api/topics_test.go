package main

import (
	"net/http"
	"testing"

	"github.com/ovh/tat"
	"github.com/ovh/tat/api/tests"
	"github.com/stretchr/testify/assert"
)

func TestTopicCreateListAndDelete(t *testing.T) {
	tests.Init(t)
	tests.Router(t)

	tests.Handle(t, http.MethodPost, "/topic", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.Create)
	tests.Handle(t, http.MethodGet, "/topics", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.List)
	tests.Handle(t, http.MethodDelete, "/topic/*topic", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.Delete)

	client := tests.TATClient(t, tests.AdminUser)
	topic, err := client.TopicCreate(tat.TopicCreateJSON{
		Topic:       "/" + tests.RandomString(t, 10),
		Description: "this is a test",
	})

	assert.NotNil(t, topic)
	assert.NoError(t, err)
	if topic == nil {
		t.Fail()
		return
	}
	t.Logf("Topic %s created", topic.Topic)
	assert.NotZero(t, topic.ID)

	topics, err := client.TopicList(nil)
	assert.NotNil(t, topics)
	assert.NoError(t, err)

	t.Log("Delete all topics")
	for _, to := range topics.Topics {
		err := client.TopicDelete(to.Topic)
		assert.NoError(t, err)
	}

}

func TestListTopicsFromCache(t *testing.T) {
	tests.Init(t)
	tests.Router(t)

	tests.Handle(t, http.MethodPost, "/topic", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.Create)
	tests.Handle(t, http.MethodPut, "/topic/add/admingroup", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.AddAdminGroup)
	tests.Handle(t, http.MethodGet, "/topics", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.List)
	tests.Handle(t, http.MethodDelete, "/topic/*topic", tests.FakeAuthHandler(t, tests.AdminUser, "TAT-TEST", true, false), topicsController.Delete)

	client := tests.TATClient(t, tests.AdminUser)
	topic, err := client.TopicCreate(tat.TopicCreateJSON{
		Topic:       "/" + tests.RandomString(t, 10),
		Description: "this is a test",
	})

	assert.NotNil(t, topic)
	assert.NoError(t, err)
	if topic == nil {
		t.Fail()
		return
	}
	t.Logf("Topic %s created", topic.Topic)

	topics, err := client.TopicList(nil)
	assert.NotNil(t, topics)
	assert.NoError(t, err)

	err = client.TopicAddAdminGroup(topic.Topic, "admin_group", false)
	assert.NoError(t, err)

	topics, err = client.TopicList(nil)
	var found bool
	for _, to := range topics.Topics {
		if to.ID == topic.ID {
			assert.NotEmpty(t, to.AdminGroups)
		}
	}
	assert.True(t, found)

	err = client.TopicDelete(topic.Topic)
	assert.NoError(t, err)
}

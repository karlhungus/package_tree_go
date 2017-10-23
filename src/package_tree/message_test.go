package main

import "testing"
import "reflect"

func TestIndexMessageParsing(t *testing.T) {
	msg := NewMessage("INDEX|foo|bar,biz")
	if msg.cmd != INDEX {
		t.Error("Expected INDEX got", INDEX, msg.cmd)
	}
  if msg.pkg != "foo" {
		t.Error("Expected pkg foo, got", msg.pkg)
	}
	deps := [2]string{ "bar", "biz" }
	if !reflect.DeepEqual(msg.dep, deps[:]) {
		t.Error("Expected dep \"bar\", \"biz\", got", msg.dep)
	}
}

func TestQueryMessageParsing(t *testing.T) {
	msg := NewMessage("QUERY|foo|bar,biz")
	if msg.cmd != QUERY {
		t.Error("Expected QUERY got", QUERY, msg.cmd)
	}
  if msg.pkg != "foo" {
		t.Error("Expected pkg foo, got", msg.pkg)
	}
	deps := [2]string{ "bar", "biz" }
	if !reflect.DeepEqual(msg.dep, deps[:]) {
		t.Error("Expected dep \"bar\", \"biz\", got", msg.dep)
	}
}

func TestRemoveMessageParsing(t *testing.T) {
	msg := NewMessage("REMOVE|foo|bar,biz")
	if msg.cmd != REMOVE {
		t.Error("Expected REMOVE got", REMOVE, msg.cmd)
	}
  if msg.pkg != "foo" {
		t.Error("Expected pkg foo, got", msg.pkg)
	}
	deps := [2]string{ "bar", "biz" }
	if !reflect.DeepEqual(msg.dep, deps[:]) {
		t.Error("Expected dep \"bar\", \"biz\", got", msg.dep)
	}
}

func TestErrorMessageParsing(t *testing.T) {
	msg := NewMessage("RE|foo|bar,biz")
	if msg.cmd != ERROR {
		t.Error("Expected ERROR got", ERROR, msg.cmd)
	}
  if msg.pkg != "" {
		t.Error("Expected pkg '', got", msg.pkg)
	}
	if len(msg.dep) != 0 {
		t.Error("Expected empty, got", msg.dep)
	}

	msg = NewMessage("QUERY|")
	if msg.cmd != ERROR {
		t.Error("Expected ERROR got", ERROR, msg.cmd)
	}
  if msg.pkg != "" {
		t.Error("Expected pkg '', got", msg.pkg)
	}
	if len(msg.dep) != 0 {
		t.Error("Expected empty, got", msg.dep)
	}

	msg = NewMessage("INDEX||")
	if msg.cmd != ERROR {
		t.Error("Expected ERROR got", ERROR, msg.cmd)
	}
  if msg.pkg != "" {
		t.Error("Expected pkg '', got", msg.pkg)
	}
	if len(msg.dep) != 0 {
		t.Error("Expected empty, got", msg.dep)
	}


}

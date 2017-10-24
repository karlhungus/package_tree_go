package main

import "testing"

func TestBasics(t *testing.T) {
	msg := NewMessage("INDEX|foo|")
	pkgr := NewPackager()
	rsp := pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("QUERY|foo|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("REMOVE|foo|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("QUERY|foo|")
	rsp = pkgr.Process(msg)
	if rsp != "FAIL\n" {
		t.Error("Expected FAIL got", rsp)
	}
}

func TestRemovalWithDeps(t *testing.T) {
	msg := NewMessage("INDEX|child|")
	pkgr := NewPackager()
	rsp := pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("INDEX|parent|child")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("QUERY|child|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("QUERY|parent|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("REMOVE|child|")
	rsp = pkgr.Process(msg)
	if rsp != "FAIL\n" {
		t.Error("Expected FAIL got", rsp)
	}

	msg = NewMessage("INDEX|parent|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("REMOVE|child|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("REMOVE|parent|")
	rsp = pkgr.Process(msg)
	if rsp != "OK\n" {
		t.Error("Expected OK got", rsp)
	}

	msg = NewMessage("QUERY|parent|")
	rsp = pkgr.Process(msg)
	if rsp != "FAIL\n" {
		t.Error("Expected FAIL got", rsp)
	}
}

package singleton

import "testing"

/*
	Criteria for test
	> When no counter has been created before, a new one is created with the value 0
	> If a counter has already been created, return this instance that holds the actual count
	> If we call the method AddOne, the count must be incremented by 1
*/

func TestGetSingleton(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		// if the value is nil, the first criteria failed
		t.Error("Singleton::GetInstance: Expected a pointer to singleton, got nil")
		// if this occurred, the following tests will result in a nil pointer dereference
		// since it will try to invoke method of a nil
		t.Fail()
	}

	if counter1.AddOne() != 1 {
		// if doesn't match, the third criteria failed
		t.Errorf("Singleton::AddOne: Expected the counter to be %d, but it is %d", 1, counter1.GetCount())
	}

	counter2 := GetInstance()

	if counter2 != counter1 {
		// if doesn't match, the second criteria failed
		t.Error("Singleton::AddOne: Expected to get the same instance, but received a different one")
	}

	if counter2.AddOne() != 2 {
		// if doesn't match, the third criteria failed
		t.Errorf("Singleton::AddOne: Expected the counter to be %d, but it is %d", 1, counter2.GetCount())
	}

	if counter1.GetCount() != counter2.GetCount() {
		// if doesn't match, the second criteria failed
		// even after passing counter2 != counter1
		t.Error("Singleton::GetCount: Expected both counter to hold same value since they are supposed to hold address of the same object")
	}

}

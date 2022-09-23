# Mocker

Used to generate [Mockery](https://github.com/vektra/mockery) mocks for Go Interfaces

## Installing
```
go get github.com/eheinen/mocker
```

## Using

In the source of the project or the directory you want to generate, just run:
```
mocker
```

The Go file needs to have an Interface with functions in order to be able to generate the mocks correctly.

Example:

```
func Test_FindPerson_Success(t *testing.T) {
    mockPersonRepository := new(mocks.Repository) // Create the Mock
    mocks.MockFindPerson(&mockPersonRepository.Mock, newPerson, nil) // Add stub for Mock with expected return
    
    actualPerson, err := New(mockPersonRepository).GetPerson() // Inject the Mock
	
    mocks.AssertFindPerson(t, &mockPersonRepository.Mock, 1) // Assert the Mock called
}

func Test_CreatePerson_Success(t *testing.T) {
    mockPersonRepository := new(mocks.Repository)
    mocks.MockPostPerson(&mockPersonRepository.Mock, nil)
    
    ctx := context.WithValue(context.Background(), "Key", "Value")
    err := New(mockPersonRepository).CreatePerson(newPerson, ctx)
    
    mocks.AssertPostPerson(t, &mockPersonRepository.Mock, 1, newPerson, ctx)
}
```

The complete example is [here](https://github.com/eheinen/mocker/tree/main/example)

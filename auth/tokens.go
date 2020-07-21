package auth

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Tokens interface {
	addToken(hash int64, id int, ctx map[string]interface{}) int64
	getToken(hash int64) int64
	rmToken(hash int64) error
}

type mapToken struct {
	accessToken int64
	userId      int
	expiresIn   *time.Timer
	signAt      time.Time
	isAdmin     bool
	ctxRoute    map[string]interface{}
	lock        *sync.RWMutex
}

func (m *mapToken) SetValue(key string, val interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.ctxRoute[key] = val
}

func (m *mapToken) Value(key interface{}) interface{} {
	if key, ok := key.(string); ok {
		return m.ctxRoute[key]
	}

	return nil
}

func (m *mapToken) GetUserID() int {
	return m.userId
}

type mapTokens struct {
	expiresIn time.Duration
	tokens    map[int64]*mapToken
	lock      sync.RWMutex
}

func (m *mapTokens) addToken(hash int64, id int, ctx map[string]interface{}) int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.tokens == nil {
		m.tokens = make(map[int64]*mapToken, 0)
	}

	m.tokens[hash] = &mapToken{
		accessToken: hash,
		expiresIn: time.AfterFunc(m.expiresIn, func() {
			m.rmToken(hash)
		}),
		userId:   id,
		ctxRoute: ctx,
		signAt:   time.Now(),
		lock:     &sync.RWMutex{},
	}

	//	todo: решить вопрос про уникальность токена
	return m.tokens[hash].accessToken
}

func (m *mapTokens) getToken(hash int64) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()

	token, ok := m.tokens[hash]
	if ok {
		return token.accessToken
	}

	return -1
}

func (m *mapTokens) rmToken(hash int64) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	_, ok := m.tokens[hash]
	if !ok {
		return errors.New("not found user in active")
	}

	delete(m.tokens, hash)

	return nil
}

package cache

import (
	"github.com/titus12/formula/opt"
	"sync"
)

var _default_expression ExpressionManager

type ExpressionManager struct {
	expressionMap map[string]*ExpressionCache
	mu            sync.RWMutex
}

func init() {
	_default_expression.init()
}

type ExpressionCache struct {
	LogicalExpression *opt.LogicalExpression
	ParamNames        []string
}

func (p *ExpressionManager) init() {
	p.expressionMap = make(map[string]*ExpressionCache, 512)
}

func (p *ExpressionManager) add(key string, value *ExpressionCache) {
	defer p.mu.Unlock()
	p.mu.Lock()
	p.expressionMap[key] = value
}

func (p *ExpressionManager) get(key string) *ExpressionCache {
	defer p.mu.RUnlock()
	p.mu.RLock()
	return p.expressionMap[key]
}

func Store(option opt.Option, originalExpression string, expressionCache *ExpressionCache) {
	_default_expression.add(originalExpression, expressionCache)
}

func Restore(option opt.Option, originalExpression string) *ExpressionCache {
	return _default_expression.get(originalExpression)
}

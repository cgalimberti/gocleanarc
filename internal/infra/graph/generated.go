package graph

import (
	"bytes"
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cgalimberti/gocleanarc/20-CleanArch/internal/infra/graph/model"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct{}

type ComplexityRoot struct {
	Mutation struct {
		CreateOrder func(childComplexity int, input model.OrderInput) int
	}

	Order struct {
		ID         func(childComplexity int) int
		Price      func(childComplexity int) int
		Tax        func(childComplexity int) int
		FinalPrice func(childComplexity int) int
	}

	Query struct {
		Orders func(childComplexity int) int
	}
}

type MutationResolver interface {
	CreateOrder(ctx context.Context, input model.OrderInput) (*model.Order, error)
}

type QueryResolver interface {
	Orders(ctx context.Context) ([]*model.Order, error)
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {
	case "Mutation.createOrder":
		if e.complexity.Mutation.CreateOrder == nil {
			break
		}
		args, err := ec.field_Mutation_createOrder_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}
		return e.complexity.Mutation.CreateOrder(childComplexity, args["input"].(model.OrderInput)), true
	case "Order.id":
		if e.complexity.Order.ID == nil {
			break
		}
		return e.complexity.Order.ID(childComplexity), true
	case "Order.Price":
		if e.complexity.Order.Price == nil {
			break
		}
		return e.complexity.Order.Price(childComplexity), true
	case "Order.Tax":
		if e.complexity.Order.Tax == nil {
			break
		}
		return e.complexity.Order.Tax(childComplexity), true
	case "Order.FinalPrice":
		if e.complexity.Order.FinalPrice == nil {
			break
		}
		return e.complexity.Order.FinalPrice(childComplexity), true
	case "Query.orders":
		if e.complexity.Query.Orders == nil {
			break
		}
		return e.complexity.Query.Orders(childComplexity), true
	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) field_Mutation_createOrder_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.OrderInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNOrderInput2githubᚗcomᚋcgalimbertiᚋgocleanarcᚋ20ᚑCleanArchᚋinternalᚋinfraᚋgraphᚋmodelᚐOrderInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, []string{"Mutation"})

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createOrder":
			out.Values[i] = ec._Mutation_createOrder(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, []string{"Query"})

	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "orders":
			out.Values[i] = ec._Query_orders(ctx, field)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _Mutation_createOrder(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object:   "Mutation",
		Field:    field,
		IsMethod: true,
	})

	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Mutation_createOrder_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}

	res, err := ec.resolvers.Mutation().CreateOrder(ctx, args["input"].(model.OrderInput))
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if res == nil {
		return graphql.Null
	}
	return ec._Order(ctx, field.Selections, res)
}

func (ec *executionContext) _Query_orders(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object:   "Query",
		Field:    field,
		IsMethod: true,
	})

	res, err := ec.resolvers.Query().Orders(ctx)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if res == nil {
		return graphql.Null
	}
	return ec._Order_slice(ctx, field.Selections, res)
}

func (ec *executionContext) _Order(ctx context.Context, sel ast.SelectionSet, obj *model.Order) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, []string{"Order"})

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Order")
		case "id":
			out.Values[i] = ec._Order_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Price":
			out.Values[i] = ec._Order_Price(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "Tax":
			out.Values[i] = ec._Order_Tax(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "FinalPrice":
			out.Values[i] = ec._Order_FinalPrice(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

func (ec *executionContext) _Order_slice(ctx context.Context, sel ast.SelectionSet, objs []*model.Order) graphql.Marshaler {
	ret := make(graphql.Array, len(objs))
	for i, obj := range objs {
		ret[i] = ec._Order(ctx, sel, obj)
	}
	return ret
}

func (ec *executionContext) _Order_id(ctx context.Context, field graphql.CollectedField, obj *model.Order) graphql.Marshaler {
	return graphql.MarshalString(obj.ID)
}

func (ec *executionContext) _Order_Price(ctx context.Context, field graphql.CollectedField, obj *model.Order) graphql.Marshaler {
	return graphql.MarshalFloat(obj.Price)
}

func (ec *executionContext) _Order_Tax(ctx context.Context, field graphql.CollectedField, obj *model.Order) graphql.Marshaler {
	return graphql.MarshalFloat(obj.Tax)
}

func (ec *executionContext) _Order_FinalPrice(ctx context.Context, field graphql.CollectedField, obj *model.Order) graphql.Marshaler {
	return graphql.MarshalFloat(obj.FinalPrice)
}

func (ec *executionContext) unmarshalNOrderInput2githubᚗcomᚋcgalimbertiᚋgocleanarcᚋ20ᚑCleanArchᚋinternalᚋinfraᚋgraphᚋmodelᚐOrderInput(ctx context.Context, v interface{}) (model.OrderInput, error) {
	res, err := ec.unmarshalInputOrderInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalInputOrderInput(ctx context.Context, v interface{}) (model.OrderInput, error) {
	if v == nil {
		return model.OrderInput{}, nil
	}
	var it model.OrderInput
	asMap := v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "id":
			var err error
			it.ID, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "Price":
			var err error
			it.Price, err = graphql.UnmarshalFloat(v)
			if err != nil {
				return it, err
			}
		case "Tax":
			var err error
			it.Tax, err = graphql.UnmarshalFloat(v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

var parsedSchema = gqlparser.MustLoadSchema(&ast.Source{Name: "schema.graphqls", Input: `type Query {
  orders: [Order!]!
}

type Order {
    id: String!
    Price: Float!
    Tax: Float!
    FinalPrice: Float!
}

input OrderInput {
    id : String!
    Price: Float!
    Tax: Float!
}

type Mutation {
    createOrder(input: OrderInput): Order
}
`, BuiltIn: false})

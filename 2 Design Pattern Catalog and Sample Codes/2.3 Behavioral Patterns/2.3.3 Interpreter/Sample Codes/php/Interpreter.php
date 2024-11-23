<?php

class Context{
}

interface Expression {
    public function interpret(Context $context): int;
}

class NumberExpression implements Expression {
    private int $number;

    public function __construct(int $number) {
        $this->number = $number;
    }

    public function interpret(Context $context): int{
        return $this->number;
    }
}

class AdditionExpression implements Expression {
    private Expression $left;
    private Expression $right;

    public function __construct(Expression $left,Expression $right) {
        $this->left = $left;
        $this->right = $right;
    }

    public function interpret(Context $context): int{
        return $this->left->interpret($context) + $this->right->interpret($context);
    }
}

class MultiplicationExpression implements Expression {
    private Expression $left;
    private Expression $right;

    public function __construct(Expression $left,Expression $right) {
        $this->left = $left;
        $this->right = $right;
    }

    public function interpret(Context $context): int{
        return $this->left->interpret($context) * $this->right->interpret($context);
    }
}

class Interpreter{
    private Context $context;

    public function __construct(Context $context) {
        $this->context = $context;
    }

    public function interpret(string $expression): int{
        // Parse expression and create expression tree
        $expressionTree = $this->buildExpressionTree($expression);
         
        // Interpret expression tree
        return $expressionTree->interpret($this->context);
    }

    private function buildExpressionTree(string $expression): Expression {
        // Logic to parse expression and create expression tree
        // For simplicity, assume the expression is already parsed
        // and represented as an expression tree
        return new AdditionExpression(
            new NumberExpression(2),
            new MultiplicationExpression(
                new NumberExpression(3),
                new NumberExpression(4)
            )
        );
    }
}


// Input expression
$expression = "2 + 3 * 4";
    
// Create interpreter
$context = new Context();
$interpreter = new Interpreter($context);
    
// Interpret expression
$result = $interpreter->interpret($expression);
echo "Result: " . $result ."\n";
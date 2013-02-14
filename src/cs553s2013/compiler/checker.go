package compiler

func Checker(input string) (ok bool, error string){
    return true, "Error"
}

type TypeDeclaration interface{
//    ident() string
//    equal() string
//    arrayType() IdebArrayType
}

type EqTypeDeclarationAST struct{
    Eq string
    Ident string
    ArrayT OfArrayTypeAST
}

func (eq EqTypeDeclarationAST) IsTree() bool{
    return true
}

type IdentTypeAST struct{
  Ident string
}

func (typ IdentTypeAST) IsTree() bool{
    return true
}

type ArrayType interface{
}

type OfArrayTypeAST struct{
    Array string
    Integer string
    Of string
    Type IdentTypeAST
}

func (OfArrayTypeAST) IsTree() bool{
   return true
}

type IdentList interface{  
}

type IdentIdentListAST struct{
    Ident ([]string)
}

type VariableDeclaration interface{
}

func (VariableDeclarationAST) IsTree() bool{
   return true
}

type VariableDeclarationAST struct{
  Colon string
  Type IdentTypeAST
  Idents ([]string)  
}

type Expression interface{
}

type ExpressionAST struct{
      SimpleExpr SimpleExpression
      RelSimpleExpression ([] RelationSimpleExpression)
}

type RelationSimpleExpression struct{
      Rel Relation 
      SimpleExpr SimpleExpression
}

type Relation interface{
}

type EqRelationAST struct{
     Eq string
}

type HashRelationAST struct{
    Hash string
}

type LessThanRelationAST struct{
    LessThan string
}

type LessOrEqRelationAST struct{
    LessOrEq string
}

type GreaterThanRelationAST struct{
    GreaterThan string
}

type GreaterOrEqRelationAST struct{
    GreaterOrEq string
}

type SimpleExpression interface{
}
  
type PlusOperatorSimpleExpression struct{
     Plus string
     Terms Term
     AddOperTerm ([]AddOperatorTerm)
}

type MinusOperatorSimpleExpression struct{
     Minus string
     Terms Term
     AddOperTerm ([]AddOperatorTerm)
}

type AddOperatorTerm struct{
     Plus string
     Minus string
     Terms Term
}

type AddOperator interface{
}

type PlusAddOperatorAST struct{
      Plus string
}

type MinusAddOperatorAST struct{
      Minus string
}

type ORAddOperatorAST struct{
      OR string
}

type Term interface{
  
}

type TimesTermAST struct{
      Factored Factor
      MulOperatorFactored ([]MulOperatorFactor)
}

type MulOperatorFactor struct{
      MultOperator MulOperator
      Factored Factor
}

type MulOperator interface{
      
}

type StarOperatorAST struct{
      Star string
}

type DIVOperatorAST struct{
      Div string
}

type MODOperatorAST struct{
      Mod string
}

type ANDOperatorAST struct{
      And string
}

type Factor interface{
      
}

type IntegerFactorAST struct{
      Integer string
}

type StringFactorAST struct{
      String string
}

type TRUEFactorAST struct{
      True string
}

type FALSEFactorAST struct{
      False string
}

type Designator interface{
  
}

type DesignatorFactorAST struct{
      Designator Designator
      ActualParameters ActualParameters
}

type ExpressionFactorAST struct{
      Expression Expression
}

type TildeFactorAST struct{
      Tilde string
      Factor Factor
}

type IdentDesignatorAST struct{
      Ident string
      Selector Selector
}

type ExpList interface{
}

type ExprExpListAST struct{
      Expression Expression
      CommaExpression ([]CommaExpression)
}

type CommaExpression struct{
      comma string
      Expression Expression
}

type Selector interface{
}

type ExpressionSelector struct{
      Expression Expression
}

type ActualParameters interface{
}

type ExpListParameters struct{
      ExpList ExpList
}

type Statement interface{
}

type AssignOrCallStatement struct{
    AssignOrCall AssignOrCall
}

type IfStatementStatement struct{
    IfStatement IfStatement
}

type WhileStatementStatement struct{
    WhileStatement WhileStatement
}

type ForStatementStatement struct{
    ForStatement ForStatement
}

type AssignOrCall interface{
}

type StatementSequence interface{
}

type StatementStatementSequence struct{
      Statement Statement
      ColonStatement ([]ColonStatement)
}

type ColonStatement struct{
      Colon string
      Statement Statement
}

type IfStatement interface{
}

type IfStatementAST struct{
    If string
    Expression Expression
    Then string
    Statements StatementSequence
    Elsif ElsifExpression
    Else ElseExpression
    End string
}

type ElsifExpression struct{
    Elsif string
    Expression Expression
    Then string
    Statements StatementSequence
}

type ElseExpression struct{
    Else string
    Statements StatementSequence
}

type ForStatement interface{
}

type ForStatementAST struct{
    For string
    Assign string
    Expression Expression
    To string
    ByInteger ByInteger
    Do string
    Statements StatementSequence
    End string
}

type ByInteger struct{
    By string
    Integer int
}

type WhileStatement interface{
}

type WhileStatementAST struct{
    While string
    Expression Expression
    Do string
    Statements StatementSequence
    End string
}

type ProcedureDeclarationAST struct{
    Heading ProcedureHeadingAST
    Colon string
    Body ProcedureBody
    Ident string
}

func (proc ProcedureDeclarationAST) IsTree() bool{
      return true
}

type SemiColonProcedureDeclaration struct{
  
}

type ProcedureBody interface{
}

type ProcedureBodyAST struct{
    DeclarationSequence DeclarationSequence
    Begin string
    StatementSequence StatementSequence
    Return string
    Expression Expression
}

type ProcedureHeading interface{
}

type ProcedureHeadingAST struct{
    Procedure string
    Ident string
    FPSectionList []FPSectionAST
    ColonIdentifier ColonIdent
}

func (ph ProcedureHeadingAST) IsTree() bool{
    return true
}

type FPSectionAST struct{
     VAR string
     Idents []string
     FT FormalTypeAST
}

func (fp FPSectionAST) IsTree() bool{
     return true
}

type FormalTypeAST struct{
     Array string
     Of string
     Ident string
}

func (ft FormalTypeAST) IsTree() bool{
  return true
}

type ColonIdent struct{
    Colon string
    Ident string
}

type DeclarationSequence interface{
}

type DeclarationSequenceAST struct{
     class string
     Type string
     TypeDeclaration ([]EqTypeDeclarationAST)
     VAR string
     VariableDeclaration ([]VariableDeclarationAST)
     ProcedureDeclaration ([]ProcedureDeclarationAST)
}

func (ds DeclarationSequenceAST) IsTree() bool{
    return true
}

func (ds DeclarationSequenceAST) Class(){
    ds.class = "DeclarationSequence"
}

type ProcedureDeclarationSemiColon struct{
     ProcedureDeclaration ProcedureDeclarationAST
     SemiColon string
}

type CommaIdent struct{
      Comma string
      Ident string
}

type FormalType interface{
}

type FT struct{
    Array string
    Of string
    ident string
}

type module interface{
  
}

type context struct{
  value string
}

type AST struct{
    Niche string
    Module string
    Ident string
    Ds DeclarationSequenceAST
    Begin string
    Ss StatementSequence
    End string
    Ident_ string
}

func (tree AST) IsTree() bool{
    return true
}

func (tree AST) Class() {
    tree.Niche = "ModuleAST"
    
}

type AbstractSyntaxTree interface{
    IsTree() bool
}

func Append(ast AbstractSyntaxTree) (success bool) {
   return true
}
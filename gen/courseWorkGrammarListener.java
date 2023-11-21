// Generated from /Users/ilyaantonov/Downloads/3 курс/ТЯП/cwp/courseWorkGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link courseWorkGrammarParser}.
 */
public interface courseWorkGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(courseWorkGrammarParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(courseWorkGrammarParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#progDeclare}.
	 * @param ctx the parse tree
	 */
	void enterProgDeclare(courseWorkGrammarParser.ProgDeclareContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#progDeclare}.
	 * @param ctx the parse tree
	 */
	void exitProgDeclare(courseWorkGrammarParser.ProgDeclareContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#varDeclare}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclare(courseWorkGrammarParser.VarDeclareContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#varDeclare}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclare(courseWorkGrammarParser.VarDeclareContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#varList}.
	 * @param ctx the parse tree
	 */
	void enterVarList(courseWorkGrammarParser.VarListContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#varList}.
	 * @param ctx the parse tree
	 */
	void exitVarList(courseWorkGrammarParser.VarListContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#programDescription}.
	 * @param ctx the parse tree
	 */
	void enterProgramDescription(courseWorkGrammarParser.ProgramDescriptionContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#programDescription}.
	 * @param ctx the parse tree
	 */
	void exitProgramDescription(courseWorkGrammarParser.ProgramDescriptionContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#assign}.
	 * @param ctx the parse tree
	 */
	void enterAssign(courseWorkGrammarParser.AssignContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#assign}.
	 * @param ctx the parse tree
	 */
	void exitAssign(courseWorkGrammarParser.AssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(courseWorkGrammarParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(courseWorkGrammarParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#subexpression}.
	 * @param ctx the parse tree
	 */
	void enterSubexpression(courseWorkGrammarParser.SubexpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#subexpression}.
	 * @param ctx the parse tree
	 */
	void exitSubexpression(courseWorkGrammarParser.SubexpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOperator(courseWorkGrammarParser.UnaryOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#unaryOperator}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOperator(courseWorkGrammarParser.UnaryOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#binaryOperator}.
	 * @param ctx the parse tree
	 */
	void enterBinaryOperator(courseWorkGrammarParser.BinaryOperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#binaryOperator}.
	 * @param ctx the parse tree
	 */
	void exitBinaryOperator(courseWorkGrammarParser.BinaryOperatorContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperand(courseWorkGrammarParser.OperandContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperand(courseWorkGrammarParser.OperandContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#id}.
	 * @param ctx the parse tree
	 */
	void enterId(courseWorkGrammarParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#id}.
	 * @param ctx the parse tree
	 */
	void exitId(courseWorkGrammarParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#const}.
	 * @param ctx the parse tree
	 */
	void enterConst(courseWorkGrammarParser.ConstContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#const}.
	 * @param ctx the parse tree
	 */
	void exitConst(courseWorkGrammarParser.ConstContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoop(courseWorkGrammarParser.LoopContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoop(courseWorkGrammarParser.LoopContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#read}.
	 * @param ctx the parse tree
	 */
	void enterRead(courseWorkGrammarParser.ReadContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#read}.
	 * @param ctx the parse tree
	 */
	void exitRead(courseWorkGrammarParser.ReadContext ctx);
	/**
	 * Enter a parse tree produced by {@link courseWorkGrammarParser#write}.
	 * @param ctx the parse tree
	 */
	void enterWrite(courseWorkGrammarParser.WriteContext ctx);
	/**
	 * Exit a parse tree produced by {@link courseWorkGrammarParser#write}.
	 * @param ctx the parse tree
	 */
	void exitWrite(courseWorkGrammarParser.WriteContext ctx);
}
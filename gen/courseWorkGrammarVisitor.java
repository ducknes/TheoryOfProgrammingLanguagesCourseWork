// Generated from /Users/ilyaantonov/Downloads/3 курс/ТЯП/cwp/courseWorkGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link courseWorkGrammarParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface courseWorkGrammarVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#program}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProgram(courseWorkGrammarParser.ProgramContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#progDeclare}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProgDeclare(courseWorkGrammarParser.ProgDeclareContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#varDeclare}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarDeclare(courseWorkGrammarParser.VarDeclareContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#varList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVarList(courseWorkGrammarParser.VarListContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#programDescription}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProgramDescription(courseWorkGrammarParser.ProgramDescriptionContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#assign}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssign(courseWorkGrammarParser.AssignContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(courseWorkGrammarParser.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#subexpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSubexpression(courseWorkGrammarParser.SubexpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#unaryOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnaryOperator(courseWorkGrammarParser.UnaryOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#binaryOperator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinaryOperator(courseWorkGrammarParser.BinaryOperatorContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#operand}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperand(courseWorkGrammarParser.OperandContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#id}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitId(courseWorkGrammarParser.IdContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#const}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConst(courseWorkGrammarParser.ConstContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoop(courseWorkGrammarParser.LoopContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#read}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRead(courseWorkGrammarParser.ReadContext ctx);
	/**
	 * Visit a parse tree produced by {@link courseWorkGrammarParser#write}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWrite(courseWorkGrammarParser.WriteContext ctx);
}
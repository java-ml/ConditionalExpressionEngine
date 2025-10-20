package com.example;

import java.util.regex.Pattern;
import java.util.regex.Matcher;

public class EscapeUtils {
    
    // 转义序列映射
    private static final String[][] ESCAPE_SEQUENCES = {
        {"\\\\", "\\"},    // 反斜杠
        {"\\\"", "\""},    // 双引号
        {"\\b", "\b"},     // 退格
        {"\\f", "\f"},     // 换页
        {"\\n", "\n"},     // 换行
        {"\\r", "\r"},     // 回车
        {"\\t", "\t"},     // 制表符
    };
    
    // Unicode转义模式
    private static final Pattern UNICODE_PATTERN = Pattern.compile("\\\\u([0-9a-fA-F]{4})");
    
    /**
     * 解码字符串中的转义字符
     */
    public static String unescapeString(String str) {
        if (str == null || str.isEmpty()) {
            return str;
        }
        
        // 移除外部的引号（如果有）
        String content = str;
        if (str.startsWith("\"") && str.endsWith("\"")) {
            content = str.substring(1, str.length() - 1);
        }
        
        // 处理Unicode转义
        content = unescapeUnicode(content);
        
        // 处理其他转义序列
        for (String[] escape : ESCAPE_SEQUENCES) {
            content = content.replace(escape[0], escape[1]);
        }
        
        return content;
    }
    
    /**
     * 解码Unicode转义序列
     */
    private static String unescapeUnicode(String str) {
        Matcher matcher = UNICODE_PATTERN.matcher(str);
        StringBuffer result = new StringBuffer();
        
        while (matcher.find()) {
            String hex = matcher.group(1);
            char unicodeChar = (char) Integer.parseInt(hex, 16);
            matcher.appendReplacement(result, Character.toString(unicodeChar));
        }
        matcher.appendTail(result);
        
        return result.toString();
    }
    
    /**
     * 编码字符串为带转义字符的形式
     */
    public static String escapeString(String str) {
        if (str == null) {
            return "\"\"";
        }
        
        StringBuilder sb = new StringBuilder();
        sb.append('"');
        
        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            switch (c) {
                case '\\': sb.append("\\\\"); break;
                case '"': sb.append("\\\""); break;
                case '\b': sb.append("\\b"); break;
                case '\f': sb.append("\\f"); break;
                case '\n': sb.append("\\n"); break;
                case '\r': sb.append("\\r"); break;
                case '\t': sb.append("\\t"); break;
                default:
                    if (c < 32 || c > 126) {
                        // 非打印字符或非ASCII字符使用Unicode转义
                        sb.append(String.format("\\u%04x", (int) c));
                    } else {
                        sb.append(c);
                    }
                    break;
            }
        }
        
        sb.append('"');
        return sb.toString();
    }
    
    /**
     * 解码标识符中的转义字符（反引号包围的标识符）
     */
    public static String unescapeIdentifier(String identifier) {
        if (identifier == null || identifier.isEmpty()) {
            return identifier;
        }
        
        // 如果是反引号包围的标识符
        if (identifier.startsWith("`") && identifier.endsWith("`")) {
            String content = identifier.substring(1, identifier.length() - 1);
            // 处理标识符中的转义
            return content.replace("\\`", "`").replace("\\\\", "\\");
        }
        
        return identifier;
    }
    
    /**
     * 编码标识符为带反引号的形式（如果需要）
     */
    public static String escapeIdentifier(String identifier) {
        if (identifier == null || identifier.isEmpty()) {
            return "``";
        }
        
        // 检查标识符是否需要转义
        if (identifier.matches("[a-zA-Z_][a-zA-Z0-9_]*")) {
            // 常规标识符，不需要转义
            return identifier;
        } else {
            // 特殊标识符，需要用反引号包围并转义
            StringBuilder sb = new StringBuilder();
            sb.append('`');
            for (int i = 0; i < identifier.length(); i++) {
                char c = identifier.charAt(i);
                if (c == '`' || c == '\\') {
                    sb.append('\\');
                }
                sb.append(c);
            }
            sb.append('`');
            return sb.toString();
        }
    }
    
    /**
     * 检查字符串是否包含需要转义的字符
     */
    public static boolean containsSpecialChars(String str) {
        if (str == null) return false;
        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            if (c == '\\' || c == '"' || c == '\b' || c == '\f' || 
                c == '\n' || c == '\r' || c == '\t' || c < 32 || c > 126) {
                return true;
            }
        }
        return false;
    }
}

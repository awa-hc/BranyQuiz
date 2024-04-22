//
//  DividerText.swift
//  BrainyQuiz
//
//  Created by Horacio Cortez on 16/04/2024.
//

import Foundation
import SwiftUI

struct DividerText: View {
    var text: String
    var color: Color
    
    var body: some View {
        HStack(spacing: 0){
            Rectangle().fill(color).frame(height: 1)
            Text(text).font(.system(size: 12, weight: .bold))
                .foregroundStyle(color).padding(.horizontal, 8)
            Rectangle().fill(color).frame(height: 1)
        }
    }
    
}

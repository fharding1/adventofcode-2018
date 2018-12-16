(defun power-level (x y serial-no)
    "Calculate the power level of a given coordinate with a given serial number"
    (- (floor (mod (/ (* (+ x 10) (+ serial-no (* (+ x 10) y))) 100) 10)) 5))

(defun make-grid (w h serial-no)
    "Creates a grid of power-levels with length h and width w"
    (loop for y from 1 to h
        collect (loop for x from 1 to w
            collect (power-level x y serial-no))))

(defun sub-grid-sum (grid size x y)
    "Sum the values of a sub-grid square with size size at 1-indexed coordinates (x, y) within a grid"
    (loop for sy from y to (+ y size -1)
          sum (loop for sx from x to (+ x size -1)
                sum (nth (1- sx) (nth (1- sy) grid)))))

;; This is just the part 1 solution because clisp is too slow to bruteforce this problem.
;; You could, however, use a Summed-area table: https://en.wikipedia.org/wiki/Summed-area_table.
(let ((grid (make-grid 300 300 9110)) (size 3) (mp 0) (mpx 0) (mpy 0))
    (loop for y from 1 to (- 300 size)
                    do (loop for x from 1 to (- 300 size)
                        do (let ((sum (sub-grid-sum grid size x y)))
                            (if (> sum mp)
                                (setq mp sum
                                      mpx x
                                      mpy y)))))
    (format t "(~a, ~a)~%" mpx mpy))